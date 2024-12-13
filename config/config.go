package config

type Config struct {
	app    App
	logger Logger
}

func Parse() (*Config, error) {
	return &Config{
		app: App{
			database: Database{
				database:       "postgresql",
				host:           "localhost",
				port:           4000,
				name:           "name",
				user:           "user",
				password:       "password",
				migrationsPath: "migrations/postgresql",
			},
			refreshToken: Token{
				secretKey: "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDAkNqABG/Xmx69\n0Vf3SvNiV88ycg3pdu/wdESeTQ5nzlGDXQAYfkcJD/tEhPBjUPJih1vNs6y05Fpx\nX6u8uBo4pVAALQs2yt8dHL0Fu7eWGAlQ1fWrORGjhiBbLo/w34gnWnz/bzJHSgpX\nZAfF8+rBZBHnc2sxYA0bzYWK0NUQspKaBO11gWp8lii6ZVvs5OHYyQBa/wW+3xj2\n0n8aAjJGYqg98J5u8AveDLCfze0hCKbHeNRw/X8cslI6bD4zio+VAGB5UcMEaWW7\nlIn44VI7uNOp+xAUS8gIM3+Q0dkLI5AD57mV44XmorgFAzUyfk1VAN5I2ynJCq7T\nw32ZLWqrAgMBAAECggEAQ4xMaS2lQd9TCEi/76ol9/BPaSWjBatH/XP7ggALtpQj\n+c2MtArPxjkJ7f4oCN6jN7a9mEmXopibtTtXFcsKv6YViwpB7cV6QE8Na05u1sK1\nudpj490P9+mbvgB2U4tFynAuere5zlIFkKQxKdonFhGRSjwtdNN5pbARUSFyK/mr\ntxCcRWhsQBcpuZKKDmfOIZLsP6VVmeqc8MtIiVFJOxBdrMrco/QnLGL8BOi8d6bN\npK88uiXRyy63179Hdy4hJVlVAIMHN8rCNQkvx6isAahxvbVunl0kyqvETUpj0czR\nZm7ttePGxRCV8uNhgls4WJ1mG98R3io4EqpEJHiNUQKBgQDwp7soOWupTgU9XILM\nYTzoAOeen4RGkAf6NNgYH5D3ZhoXoHTCXdchB37dZbDDGD4NFMRbg13HUNyYFPeP\n6m7N9n07rXo512KWOfEWuihNK1zUrInMDn4tKD9aVwpmeIaDpQOQNM7aspjLxGLe\nACvcStOOlyxU8RlcQeXffGmlCQKBgQDM2CYfAtEyDzQEC+lm96nXPYE3Ht7r3jQx\nxegUi3Fpxqog0k5PPvqdknq2LMxb4oPI6IOo5ZPqKcrvsWieT9EQrYG/c5GEZpC6\nGxLga/ZZH8Qnz9973BGnZLrz06rnxkPhzpclJk/ZFZnL9u/09baOLDjLnBxpA2pr\n2Wi//YKTEwKBgEgPcZNR3lML8bW2BdR9KWevHav5x/n5Fups+GgAy7/I6vc4LL4m\nq0xQmkRhNtS0ZaA1H7bnVyVfqS5A2QBXCmbR8B32jCrVu8eQAgj0JS02Vs23/GsF\ny/Jpos7CUXLtlBnHR5HevjHuWKJrY89l7bxBC6gezD+DEjz13odVEa1BAoGAAWJ4\np38CYP8/T8Ow2TRVXpIY3IGF6jw/SBLZZcInixp/receBmUblpMV50NYQkq+r037\nDdo3bL7Jk2M0nNVXSXCSzvN+CVIUKl3ie8ffC20ui0JZdWHgcKw6o0lL/3MUpGMc\nHHEs2v/NYVaS+qnqiP19iMrOWltejcoTbjDNgCcCgYEAqXJB7aVp8CqDw5nWhbuV\ndqO0f+7YSSqVdd+1xZrHdb4PXojV/OeiN51UAPW6g7U+daiMLa8PtBJNxz8csRyf\nI5l/wr9sVVTeu62x8W7afwLIsndCtf8plW3R9GIh3hrbS/TScFBwSyxwlry4Y/CS\ndNcn/GPJ0PDXJbbVnRgYBZc=\n-----END PRIVATE KEY-----",
				ttlSecond: 2592000,
			},
			accessToken: Token{
				secretKey: "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDG/D22OsvHBdQY\nuQjGZJiY7ixozC9A0McLH7HI2AeEwp+/qkU77SkEQXmG7LbA6/NHkDmA3qE3UPKO\nN+W1nv/JR5yQE+YN2ulDwWkn9dzmwoHS2Nik7N6tobdz4zt0rAmtrHWvUa7b9vmL\nDWcuxsXUIKGhU8LguwKyvCOSJnNAFgEBOJYU5PoRSzw+EqLviRb7DNMz1lva6KQ3\nD4S9TviQYgWkunYhlwobtPGXJwENuDzAYsv4ThBaZPwiWyVdXea9xrXSsc+g2044\nUaj30VL2DVePvYOqz1vCi50WEJiL6xj7cX9KyFWO7zx4SLxhKaAQ/tfGVve+W+V2\nZ+4pkddDAgMBAAECggEAUFB3SBorLcV4ehIVLJ5lyTQ40IO5ftvFuwOr6njnb/pu\nKBo9n5Z1u7YBqyfYC74wj23zB8TUJVOS0zxUhx4J99/jnXOqo2T8SdDC9NITu0A+\nzi+gIdFJf6OKPEg3MzxNG8BEL+M0RKxGxI1AXGz2ZExXcJoySiqxSBsP8Gl8kxyX\nt4KWIFccHqPdMWagwAeoNqwvfW4NCZ5atvBGHZLEl5vxdMORFNdjscayuhHUULGM\nWd/5/Egk/A1OKwpQp1dmpScwQcvqV98N1wAzLUbwDLBLCW2+LjF20quPXlDQIG4n\n9lHmQhj0p/C1KrRSOhxDor/57GYIqgYF3nDxsji6QQKBgQD0xO9eMYEen1EfkMoa\nO3oLmKxEmNLAi+6Yra8vQX5SJrBehobtA+ldy7s13dv6NVMvsQorhYCMsuqmKAbd\nLFOgW1nPbdfsBXVh6SBHfstUG/473OQghy40PMDokyPU15IC6HoEDV503kT1KSir\nRHKT5ovy/TvyYVBth9KfKN5SiQKBgQDQHYbU1egNj1IRAy+pCmG660H0NTnvFWse\nVoTeFEC864cPSLxJMvcaly6H9Xv5T34McjZoMqz799IwlD3yxxX8PxBdk9hqrIir\naKU5PLCzcfXLH4FSdIBk/tHsqfHF+ffWvLzntbynvJM8As9r5/c26boKBcRfwSpw\nVwnoHYSYawKBgBXf98iCQhjbgiOBDB7WC+03YsXqpoMgvl0UwKBHiSOMY5y/BxXt\nkJCaYYvJ3Rv6YgnUF6WWE8v/PWtGDn77tCmToG4iUfa88iWi8QiSs1c7/TeaYSTr\nD90JRPBEHiKqdVkfL7BnBXcKWj1SxVixEFBZfU5lBVN/EPII/nDM9TmJAoGACRUA\nyVD8a3sRrcRh1BKr6ShTHIEmIhThM9W6vqZLnPL6VeRqsxcSSYrHcQ95dYvOiGSD\nk1CC1AwglBYkDfA6OWPXJv5It38IB9LtQu4vF+WlQFQDpMGIyN+F8boI3wblle1u\noP4BWaCVHE1TF/Zrh23zvHQ7+aHIZFaoDAUdwDkCgYBOEXVuwu0bNdPAoqnWIvUQ\n+9AmW1kB/P7wnNwB89jgX705GPD7aBQrkv4IQgEbWR/wuJJthAym1VasHhLsNG3o\nF0pQkdguhNlBwriaB4ntVfdREw5ok9KK2J0EvW0DMoQK9tvjfChP5roXFdEZP7mn\nOvFZ6Z+Pws+/ko6T0mng4A==\n-----END PRIVATE KEY-----",
				ttlSecond: 1800,
			},
			grpc: GRPCServer{
				port: 6000,
			},
		},
		logger: Logger{
			level: "debug",
		},
	}, nil
}

func (c Config) App() App {
	return c.app
}

func (c Config) Logger() Logger {
	return c.logger
}
