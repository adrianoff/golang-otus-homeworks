package hw10programoptimization

import (
	"bufio"
	"errors"
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
			splitN := strings.SplitN(user.Email, "@", 2)
			if len(splitN) < 2 {
				return nil, errors.New("invalid email: " + user.Email)
			}

			curDomain := strings.ToLower(splitN[1])
			result[curDomain]++
		}
	}
	return result, nil
}
