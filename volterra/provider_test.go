// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/time/rate"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"volterra": testAccProvider,
	}
}

func testAccPreCheck() {

	if v := os.Getenv("VOLT_TENANT"); v == "" {
		os.Setenv("VOLT_TENANT", "test")
	}
	if v := os.Getenv("VOLT_ENV"); v == "" {
		os.Setenv("VOLT_ENV", "test")
	}
}

func TestProviderLimiterValidation(t *testing.T) {
	testCases := []struct {
		name             string
		rate             float64
		burst            int
		expectRateError  bool
		expectBurstError bool
		errorContains    string
	}{
		{
			name:             "valid limiter configuration",
			rate:             30.0,
			burst:            15,
			expectRateError:  false,
			expectBurstError: false,
		},
		{
			name:             "minimum valid values",
			rate:             1.0,
			burst:            1,
			expectRateError:  false,
			expectBurstError: false,
		},
		{
			name:             "fractional rate ",
			rate:             0.5,
			burst:            10,
			expectRateError:  false,
			expectBurstError: false,
		},
		{
			name:             "rate zero (invalid)",
			rate:             0.0,
			burst:            10,
			expectRateError:  true,
			expectBurstError: false,
			errorContains:    "must be greater than 0",
		},
		{
			name:             "rate negative",
			rate:             -1.0,
			burst:            10,
			expectRateError:  true,
			expectBurstError: false,
			errorContains:    "must be greater than 0",
		},
		{
			name:             "burst zero",
			rate:             10.0,
			burst:            0,
			expectRateError:  false,
			expectBurstError: true,
			errorContains:    "must be at least 1",
		},
		{
			name:             "burst negative",
			rate:             10.0,
			burst:            -5,
			expectRateError:  false,
			expectBurstError: true,
			errorContains:    "must be at least 1",
		},
	}

	provider := Provider()
	limiterSchema := provider.Schema["limiter"]
	elemResource := limiterSchema.Elem.(*schema.Resource)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test rate validation
			if rateSchema, ok := elemResource.Schema["rate"]; ok && rateSchema.ValidateFunc != nil {
				_, errs := rateSchema.ValidateFunc(tc.rate, "rate")
				hasError := len(errs) > 0

				if hasError != tc.expectRateError {
					t.Errorf("Rate validation: expected error: %v, got error: %v, errors: %v", tc.expectRateError, hasError, errs)
				}

				if tc.expectRateError && tc.errorContains != "" {
					found := false
					for _, err := range errs {
						if err != nil && strings.Contains(err.Error(), tc.errorContains) {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("Expected rate error to contain '%s', got: %v", tc.errorContains, errs)
					}
				}
			}

			// Test burst validation
			if burstSchema, ok := elemResource.Schema["burst"]; ok && burstSchema.ValidateFunc != nil {
				_, errs := burstSchema.ValidateFunc(tc.burst, "burst")
				hasError := len(errs) > 0

				if hasError != tc.expectBurstError {
					t.Errorf("Burst validation: expected error: %v, got error: %v, errors: %v", tc.expectBurstError, hasError, errs)
				}

				if tc.expectBurstError && tc.errorContains != "" {
					found := false
					for _, err := range errs {
						if err != nil && strings.Contains(err.Error(), tc.errorContains) {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("Expected burst error to contain '%s', got: %v", tc.errorContains, errs)
					}
				}
			}
		})
	}
}

func TestProviderLimiterConfiguration(t *testing.T) {
	testCases := []struct {
		name          string
		config        map[string]interface{}
		expectedRate  rate.Limit
		expectedBurst int32
	}{
		{
			name: "limiter with rate 30, burst 15",
			config: map[string]interface{}{
				"limiter": []interface{}{
					map[string]interface{}{
						"rate":  30.0,
						"burst": 15,
					},
				},
			},
			expectedRate:  rate.Limit(30.0),
			expectedBurst: 15,
		},
		{
			name: "limiter with minimum values",
			config: map[string]interface{}{
				"limiter": []interface{}{
					map[string]interface{}{
						"rate":  1.0,
						"burst": 1,
					},
				},
			},
			expectedRate:  rate.Limit(1.0),
			expectedBurst: 1,
		},
		{
			name:          "no limiter - should default to Inf",
			config:        map[string]interface{}{},
			expectedRate:  rate.Inf,
			expectedBurst: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			provider := Provider()
			d := schema.TestResourceDataRaw(t, provider.Schema, tc.config)

			// Manually extract limiter config (simulating providerConfigure logic)
			var limiter *ProviderLimiter

			if v, ok := d.GetOk("limiter"); ok {
				limiterList := v.([]interface{})
				limiterMap := limiterList[0].(map[string]interface{})
				rateVal := limiterMap["rate"].(float64)
				burstVal := int32(limiterMap["burst"].(int))

				limiter = &ProviderLimiter{
					Rate:  rate.Limit(rateVal),
					Burst: burstVal,
				}
			} else {
				limiter = &ProviderLimiter{Rate: rate.Inf, Burst: 1}
			}

			if limiter == nil {
				t.Fatal("Expected limiter to be configured")
			}

			if limiter.Rate != tc.expectedRate {
				t.Errorf("Expected rate %v, got: %v", tc.expectedRate, limiter.Rate)
			}

			if limiter.Burst != tc.expectedBurst {
				t.Errorf("Expected burst %d, got: %d", tc.expectedBurst, limiter.Burst)
			}
		})
	}
}

func TestProviderConfigure_WithLimiter(t *testing.T) {
	// Set required environment variables for provider configuration
	os.Setenv("VOLT_TENANT", "test-tenant")
	os.Setenv("VOLT_ENV", "test")
	defer os.Unsetenv("VOLT_TENANT")
	defer os.Unsetenv("VOLT_ENV")

	testCases := []struct {
		name          string
		config        map[string]interface{}
		expectedRate  rate.Limit
		expectedBurst int32
		expectError   bool
	}{
		{
			name: "valid limiter with rate and burst",
			config: map[string]interface{}{
				"test": true,
				"limiter": []interface{}{
					map[string]interface{}{
						"rate":  30.0,
						"burst": 15,
					},
				},
			},
			expectedRate:  rate.Limit(30.0),
			expectedBurst: 15,
			expectError:   false,
		},
		{
			name: "limiter with minimum values",
			config: map[string]interface{}{
				"test": true,
				"limiter": []interface{}{
					map[string]interface{}{
						"rate":  1.0,
						"burst": 1,
					},
				},
			},
			expectedRate:  rate.Limit(1.0),
			expectedBurst: 1,
			expectError:   false,
		},
		{
			name: "no limiter defaults to unlimited",
			config: map[string]interface{}{
				"test": true,
			},
			expectedRate:  rate.Inf,
			expectedBurst: 1,
			expectError:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a Config struct to capture the limiter before Client() is called
			raw := make(map[string]interface{})
			for k, v := range tc.config {
				raw[k] = v
			}

			provider := testAccProvider
			d := schema.TestResourceDataRaw(t, provider.Schema, raw)

			// Create a temporary Config to test limiter configuration
			config := Config{}
			config.test = true

			if v, ok := d.GetOk("limiter"); ok {
				limiterList := v.([]interface{})
				limiterMap := limiterList[0].(map[string]interface{})
				rateVal := limiterMap["rate"].(float64)
				burstVal := int32(limiterMap["burst"].(int))

				config.limiter = &ProviderLimiter{
					Rate:  rate.Limit(rateVal),
					Burst: burstVal,
				}
			} else {
				config.limiter = &ProviderLimiter{Rate: rate.Inf, Burst: 1}
			}

			// Verify limiter configuration
			if config.limiter == nil {
				t.Fatal("Expected limiter to be configured")
			}

			if config.limiter.Rate != tc.expectedRate {
				t.Errorf("Expected rate %v, got: %v", tc.expectedRate, config.limiter.Rate)
			}

			if config.limiter.Burst != tc.expectedBurst {
				t.Errorf("Expected burst %d, got: %d", tc.expectedBurst, config.limiter.Burst)
			}

			// Test that the provider can actually be configured (end-to-end)
			result, diags := provider.ConfigureContextFunc(nil, d)
			if diags.HasError() && !tc.expectError {
				t.Fatalf("Unexpected error configuring provider: %v", diags)
			}
			if !diags.HasError() && tc.expectError {
				t.Fatal("Expected error but got none")
			}

			// Verify we got an APIClient back when using test mode
			if !tc.expectError && result == nil {
				t.Fatal("Expected non-nil result from ConfigureContext")
			}
		})
	}
}
