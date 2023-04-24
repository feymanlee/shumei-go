/**
 * @Author: lifameng@changba.com
 * @Description:
 * @File:  client_options
 * @Date: 2023/4/3 15:14
 */

package shumei

import "time"

// ClientOption are options that can be passed when creating a new Client
type ClientOption func(*Client) error

// WithTimeout is a Client option that allows you to override the default timeout duration of requests
// for the Client. The default is 30 seconds. If you are overriding the http Client as well, just include
// the timeout there.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) error {
		c.timeout = timeout
		return nil
	}
}

//
// WithRegion
//  @Description:
//  @param region
//  @return ClientOption
//
func WithRegion(region string) ClientOption {
	return func(c *Client) error {
		c.region = region
		return nil
	}
}
