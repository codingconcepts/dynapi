package dynapi

import "fmt"

// Option allows for server configuration to creation time.
type Option func(s *Server) error

// Host allows for the configuration of the TLS host name.
func Host(value string) Option {
	return func(s *Server) error {
		s.host = value
		return nil
	}
}

// Port allows for the configuration of a port number.
func Port(value int) Option {
	return func(s *Server) error {
		s.port = value
		return nil
	}
}

// CertsDir allows for the configuration of a certificates
// directory.
func CertsDir(value string) Option {
	return func(s *Server) error {
		s.certsDir = value
		return nil
	}
}

// BuildInfo allows for the configuration of a build-time
// information to store in the server.
func BuildInfo(version string, timeStamp string) Option {
	return func(s *Server) error {
		s.buildVersion = version
		s.buildTimestamp = timeStamp
		return nil
	}
}

// Routes allows for the configuration of route endpoints.
func Routes(values ...RouteConfig) Option {
	fmt.Println(values)
	return func(s *Server) error {
		for _, value := range values {
			fmt.Println(value)
			s.add(value)
		}
		return nil
	}
}
