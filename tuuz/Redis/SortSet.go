package Redis

import "github.com/go-redis/redis/v8"

func SortSet_Add(key string, score float64, value interface{}) error {
	var z redis.Z
	z.Score = score
	z.Member = value
	return goredis.ZAdd(goredis_ctx, key, &z).Err()
}

func SortSet_Count(key string, min, max interface{}) int64 {
	if min == nil || max == nil {
		count, err := goredis.ZCard(goredis_ctx, key).Result()
		if err != nil {
			return 0
		}
		return count
	} else {
		minuim, ok := min.(string)
		if !ok {
			return 0
		}
		maxium, ok := max.(string)
		if !ok {
			return 0
		}
		count, err := goredis.ZCount(goredis_ctx, key, minuim, maxium).Result()
		if err != nil {
			return 0
		}
		return count
	}
}

func SortSet_Increase(key string, incr float64, value string) (float64, error) {
	return goredis.ZIncrBy(goredis_ctx, key, incr, value).Result()
}

func SortSet_rank(key, rank_by string) (int64, error) {
	return goredis.ZRank(goredis_ctx, key, rank_by).Result()
}

func SortSet_list() {

}

func SortSet_search(key, search string, limit int) (ret []string, err error) {
	ret, _, err = goredis.ZScan(goredis_ctx, key, 0, search, int64(limit)).Result()
	return
}
