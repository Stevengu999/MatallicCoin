package visor

import "github.com/skycoin/skycoin/src/coin"

const (
	// Maximum supply of skycoins
	MaxCoinSupply uint64 = 3e8 // 100,000,000 million

	// Number of distribution addresses
	DistributionAddressesTotal uint64 = 100

	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal

	// Initial number of unlocked addresses
	InitialUnlockedCount uint64 = 100

	// Number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 5

	// Unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 60 * 60 * 24 * 365 // 1 year
)

func init() {
	if MaxCoinSupply%DistributionAddressesTotal != 0 {
		panic("MaxCoinSupply should be perfectly divisible by DistributionAddressesTotal")
	}
}

// Returns a copy of the hardcoded distribution addresses array.
// Each address has 1,000,000 coins. There are 100 addresses.
func GetDistributionAddresses() []string {
	addrs := make([]string, len(distributionAddresses))
	for i := range distributionAddresses {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// Returns distribution addresses that are unlocked, i.e. they have spendable outputs
func GetUnlockedDistributionAddresses() []string {
	// The first InitialUnlockedCount (30) addresses are unlocked by default.
	// Subsequent addresses will be unlocked at a rate of UnlockAddressRate (5) per year,
	// after the InitialUnlockedCount (30) addresses have no remaining balance.
	// The unlock timer will be enabled manually once the
	// InitialUnlockedCount (30) addresses are distributed.

	// NOTE: To have automatic unlocking, transaction verification would have
	// to be handled in visor rather than in coin.Transactions.Visor(), because
	// the coin package is agnostic to the state of the blockchain and cannot reference it.
	// Instead of automatic unlocking, we can hardcode the timestamp at which the first 30%
	// is distributed, then compute the unlocked addresses easily here.

	addrs := make([]string, InitialUnlockedCount)
	for i := range distributionAddresses[:InitialUnlockedCount] {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// Returns distribution addresses that are locked, i.e. they have unspendable outputs
func GetLockedDistributionAddresses() []string {
	// TODO -- once we reach 30% distribution, we can hardcode the
	// initial timestamp for releasing more coins
	addrs := make([]string, DistributionAddressesTotal-InitialUnlockedCount)
	for i := range distributionAddresses[InitialUnlockedCount:] {
		addrs[i] = distributionAddresses[InitialUnlockedCount+uint64(i)]
	}
	return addrs
}

// Returns true if the transaction spends locked outputs
func TransactionIsLocked(inUxs coin.UxArray) bool {
	lockedAddrs := GetLockedDistributionAddresses()
	lockedAddrsMap := make(map[string]struct{})
	for _, a := range lockedAddrs {
		lockedAddrsMap[a] = struct{}{}
	}

	for _, o := range inUxs {
		uxAddr := o.Body.Address.String()
		if _, ok := lockedAddrsMap[uxAddr]; ok {
			return true
		}
	}

	return false
}

var distributionAddresses = [DistributionAddressesTotal]string{
	"xpeEY235EsyQ1CPVWh58hhYawELAHPRsdK",
	"Fw4vsWcR3qRkCW1me5RDmY5Kju3UwuBEPz",
	"2Q4ZJsRUMSY9HiqPrNFeonjnGU2ZZDPrNGD",
	"xc6VnHYWWDcpHT8642YFPxFM7rJqR3EKoB",
	"QseStsHciezXNC5mXcsKp72hPWKGPxPtZz",
	"2RvYQqshdgTNo9r43nfaFYfkU1YNPgpTecg",
	"YHjirfYqhK7SsG3SLMw1eGWaDz2BvZJ4js",
	"CgBaDnrmoKkDJnkFu3XzEvkW28SjXZNjGs",
	"YNKHrQwanibx61D9W23Zwaxp5oQhWdANuk",
	"swpFGZvKapSyZN65rDehoA1KZMgUxTcgn2",
	"o6qSX21iP5mYWU5CJTDBcEYJgxwrGvXo2a",
	"fRVyaRWYXufbS1HY9zuuJUE8HiNMDDG9SF",
	"z8rCERbFeGWyoXqA9piSzzjbnJSUFbxSEn",
	"2ertAbHwBapUdckDymQqBhEEckHHsboaC3E",
	"BTv3E2ojitgnh1E9P1PwRUf4EHWmMAkuNK",
	"27uLqCZVt3nVnN1AUCUE8joFMJcGrU9Wjhp",
	"YzGsk8L2SvtAp23uDfHH6WC7C4LxMyzUK7",
	"2PmtxGiCKMuEBmGho958M8VxJs2ze8qH47k",
	"ChaVBzmFpPj2HN3pJeC2NxdY41QqvfBoNf",
	"2iTACbp8ZVX2L5Rti3SdViRVqk2gXfU6tvE",
	"k3HaDpktNc2XMgpeXp6WXMARVJvwfsW6iS",
	"nP3dhqAeD8EaqaTS1KCvttzoFu6u49jjD9",
	"2kVn9fw3bcusHKWtzjStNa7ASxrGarKF72Q",
	"iiw4qEixwZcJDN7NFwsWsBAEchmiNPPQqY",
	"BhLnXedLZkA3aQ9YkTSuA4ghMJ6j7xqjsD",
	"cjJgQN3qcRy9az9Ju5XERKDK14XMjzjYPd",
	"fNZCGRpsNZyNrmBNqsdJDZJuMoYqipVBmb",
	"9UP4wkSBxsSYDNWiLXkWYswhWogHXw12cN",
	"b3pbhVpNw9c5msFLaRv3yJbXRxiALpVdJi",
	"hvHev9jU4uww9NuzZfzFmg88QPRBGSTbDU",
	"2dPCpsdaKbSmkDXXDL2rmq5152sQk2vgY3B",
	"ZCpWnp1cEdZM14RpPyosizmwwkqzZ8KCfF",
	"2HFUKepZPUkPdMhTzmasZ2EDkF72gp9oNNY",
	"eXschg46oiuiRtWDpiWnNhZeWf92o38SNU",
	"qCTBXwTK4tdq1CkYMfvJezwDiUHnwiqKnJ",
	"2gJKUgXPhzJ8oRADZ8QTg94ibpKK6RXuQvT",
	"21FLaFSrWv8mf5jD6eXBy1AbNBXPRmqk4DN",
	"eZCZew1xX4WQMU7hgDEo3fMN8HD47u3L7n",
	"7vDvmHz6Qo6fuFNKCtWjZjBK4HSUfmoj2T",
	"XcjDwQn66CU2u32Lz3MGMwx7xiCMRzDRWZ",
	"7crtTMycEydj5354EatQJWL48tuBWcjreV",
	"2NYLYECy6YQM2jTccYXWMLvd974okUng5Vt",
	"2cbxDVL7Uapazr2axMAjEMBKLfnkHJkCQmt",
	"PZ2SdYt9gRmYzhJNpNmo7GGYiPPYqin9Qu",
	"2mJj5o1xKSXxS4Tr3YWMeTtWL9kXtpwS42R",
	"ikvotuTQbtTHHJB3tD6o7r2jTZ2NCa5vzh",
	"2mtqfbPxM8UCNqyGBCKn7xcxjHzKwegksGw",
	"KGyvwN12ajtwtdMz2rwqwsyF2RGTYJM8n2",
	"2bfbvuLksRtTWCwUhLg1BTTJJx8yH7XmBLy",
	"jGwDJZAs3ct5268CycGUAxwczixdkhLDiw",
	"2ZPzYfjvXFXfAAmKHEi7oq5bDKBDPABqZXo",
	"2apX9iiv849Crt1Z8bD6ZWHdrj8hPiqEyiG",
	"22SarJQZKKmX9DkW3XGwkJFxNfBfPAnfXmA",
	"2V7aLJZe7TZHHCiDjZyJchzjsHSM1S2U7c2",
	"FKFbpUVGUgbfPYa5ukCqFvvnp8BTWsNFqE",
	"uBQAXN5pBfiiAd1em3sNTrrvgP3vbhPRA3",
	"2W4EziLmuZyUXaez5i7RrdRJ5qrvi9e2Afz",
	"bMEk8iJ8wWjopTGhizaTgTNckqyJ6iAtuj",
	"2VWg5mN8Zmu2uVxTwkXyBK6EjMVtNFgirPF",
	"kjW912BzscLnt6moZPdQwcR2JJKZVB4czP",
	"24rd9vzvFtEd9HGM3iCU4xFtCkzTcs93pDm",
	"FBzVV6HG3KvPMpVXKmtQLe6EaPAe2mEb2m",
	"H6sPcivcVUaskugHuzja1cBGLtakv7Wpx6",
	"jAiN3az9n3Kf43F8NRhsyNCciuRqVBPFhY",
	"2eWXto9QnSKzuUWM4eTcufpgDWva5XH49AN",
	"sXtUh5Kj7k3X5ZAXbXr7a1NXNsWBUvS277",
	"a9tAdEdaBtu5kpsESAkPWZyR3WU7Lg9X8f",
	"2d14E4n7AtHgrrZSfx7SmMuWug8y1ohf1rU",
	"dUYj5d12SvGpWMD6Vw5wTt6YuvMiHAFBbZ",
	"2VFPsYux92bmbwuxXc9v2hXxtyB9j7icxsC",
	"24wiQ8mk7LP3SXz5fgAzBo8ypLKdaeyxaeH",
	"2caj83DQ69D7yNYshyufZyoaW9jHGfgm1Dk",
	"LNGF6WmiGUQy8bsJMbjexoBFTzKb7GwzB9",
	"2BDhU1jKSfaSRys5KLYDUNJ5fdJwRXRxEUP",
	"2f5HaWoSuPKk5Rg9XH8yqYv4Q7ZHhyg7R6W",
	"26EfnxfxGTT3iRgiubvvRGrm759AkPKvy3s",
	"jXRX8E1LZLLbrDVb1BjBU1tnwTc1KYa6VH",
	"WbvwHt9vgD2EgmZQSd3M7wTAcbvty5H1si",
	"g5uUKfXqFuqi9Wjf5FktJwj5bRk1GieP65",
	"ggnCSXgY8MW1nC2opNtYs7pSev8CmbY8Dz",
	"tc3jhcLEspMgtHZCZYKdF8RxhTZcYr5XA2",
	"vTHK2mQrZFeHH4mxELUww56dvdPmGUwypb",
	"dDjA38XdUrZNpyxk1vh7nZkuYUeVaLCFNE",
	"2RWngCZKWCk7q9vR8GPMzCnEFEuYo7obZRi",
	"2BAEdDzCLs7dTXFQhtgvhjQXHkd4YENB1pS",
	"2RZy9kTeoU4t6WjvZnC3czPLsduU3gpSKhK",
	"2bYEmu2SsQuckZJEG5JAp9P7QtewV6r4xuq",
	"22UWnvSj21uM3R6jfMYELi3hsC6UawKriYt",
	"JasfB36jxYE3egtQbXZN8VrmZyehHDUhcs",
	"Tr2B2vQ9RZA5pfqagpX4ui1o82n4eJgmgk",
	"aWH9BYvn4AkH1ZtWN98KcfzQhyyMcrX9XY",
	"2i1gfFwGpYVky9ahtrSgXgVuhfwGzVjDF5d",
	"2d2ig9Xxuina5qdqGm8JZGDBX8L1H5UNrsR",
	"2AoREsgSoSpwjUMQnpDut3wQNJb8jBUc8xf",
	"Cj1tR6z8DMhiyFGrnFsCyK59TGQfW89UUJ",
	"2WBKtwEUo1cWggM3pT7xhDQTmm6HW9KZ9su",
	"26qEhx9Nfpav5DdzJEVTZcUyb3KWBxV1yte",
	"21WmsuFoDogMGUmvwMfEf1aXs5YZHwNYM7x",
	"MqUk3JLmeuyhq9TrFC1uJDuKgYd2GGDDjB",
	"22njCVTYPKcXXFGWAHQADXowtJfKDyXHnHA",
}
