package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat)
	scanner := bufio.NewScanner(r)

	user := &User{}
	for scanner.Scan() {
		if err := user.UnmarshalJSON(scanner.Bytes()); err != nil {
			return nil, err
		}

		if strings.HasSuffix(user.Email, "."+strings.ToLower(domain)) {
			split := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			result[split]++
		}
	}
	return result, nil
}
