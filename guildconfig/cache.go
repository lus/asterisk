package guildconfig

// cache defines the guild configuration cache
var cache = make(map[string]*GuildConfig)

// RetrieveCached returns the guild cinfiguration corresponding to the given guild ID but checks the cache first
func RetrieveCached(guildID string) (*GuildConfig, error) {
	if cache[guildID] != nil {
		return cache[guildID], nil
	}

	guildConfig, err := Retrieve(guildID)
	if err != nil {
		return nil, err
	}
	cache[guildID] = guildConfig
	return guildConfig, err
}

// RemoveFromCache removes the guild configuration corresponding to the given guild ID from the cache
func RemoveFromCache(guildID string) {
	delete(cache, guildID)
}
