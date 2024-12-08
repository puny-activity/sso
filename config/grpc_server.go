package config

type GRPCServer struct {
	port int
}

func (c GRPCServer) Port() int {
	return c.port
}
