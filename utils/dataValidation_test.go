package utils

import (
    "testing"
)

func TestGetValidatedPrice(t *testing.T) {
    testCases := []struct {
        name  string
        price string
        out   string
        err   string
    }{
        {
            name:  "validPrice",
            price: "10.0",
            out:   "10.0",
        },
        {
            name:  "validPrice",
            price: "10.01",
            out:   "10.01",
        },
        {
            name:  "validPrice",
            price: "0.1",
            out:   "0.1",
        },
        {
            name:  "validPrice",
            price: "0.01",
            out:   "0.01",
        },
        {
            name:  "invalidPrice",
            price: "10.001",
            out:   "",
        },
        {
            name:  "invalidPrice",
            price: "10000000000.0",
            out:   "",
        },
        {
            name:  "invalidPrice",
            price: ".0",
            out:   "",
        },
        {
            name:  "invalidPrice",
            price: "10.",
            out:   "",
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            out, err := DataValidation.GetValidatedPrice(tc.price)
            if out != tc.out || err.Error() != tc.err {
                t.Fatalf("Dropped on test with price=%s", tc.price)
            }
        })
    }
}
