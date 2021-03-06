/*
 * Subscription Manager
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package workspaceEngine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PlansCurrentGet - Returns an instance of APlan that a logged on user is currently using.\"
func PlansCurrentGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PlansGet - Returns a list of available subscription plans.\"
func PlansGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
