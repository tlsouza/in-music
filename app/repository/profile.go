package repository

import (
	"api/app/types"
	"errors"
	"sync"
)

var lockProfileRepository = &sync.Mutex{}
var singleProfileRepositoryInstance *ProfileRepository

// InMemoryRepository struct implementing Repository interface
type ProfileRepository struct {
	profiles []types.Profile
	mu       sync.Mutex // to ensure thread-safe operations
	nextID   uint64     // auto-incrementing ID
}

func GetProfileRepositoryInstance() *ProfileRepository {
	if singleProfileRepositoryInstance == nil {
		lockProfileRepository.Lock()
		defer lockProfileRepository.Unlock()
		if singleProfileRepositoryInstance == nil {
			singleProfileRepositoryInstance = &ProfileRepository{
				profiles: []types.Profile{},
				nextID:   1,
			}
		}
	}
	return singleProfileRepositoryInstance
}

// Save adds a new profile to the repository
func (r *ProfileRepository) Save(profile types.Profile) (uint64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Set an auto-incrementing ID
	profile.ID = r.nextID
	r.nextID++

	// Add the profile to the slice
	r.profiles = append(r.profiles, profile)
	return profile.ID, nil
}

// GetByID retrieves a profile by its ID
func (r *ProfileRepository) GetByID(id uint64) (*types.Profile, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, profile := range r.profiles {
		if profile.ID == id {
			return &profile, nil
		}
	}
	return nil, errors.New("profile not found")
}

// GetAll retrieves all profiles
func (r *ProfileRepository) GetAll() []types.Profile {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.profiles
}
