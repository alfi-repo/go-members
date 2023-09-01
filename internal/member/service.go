package member

func NewService(repository Repository) Service {
	return Service{
		repository: repository,
	}
}

type Service struct {
	repository Repository
}

func (s *Service) new(r newRequest) (response, error) {
	member, err := NewMember(r.Name, r.Verified)
	if err != nil {
		return response{}, err
	}

	if err = s.repository.Save(&member); err != nil {
		return response{}, err
	}
	return toResponse(member), nil
}

func (s *Service) findByID(r findRequest) (response, error) {
	member, err := s.repository.FindByID(r.ID)
	if err != nil {
		return response{}, err
	}
	return toResponse(member), nil
}

func (s *Service) edit(r editRequest) (response, error) {
	member, err := s.repository.FindByID(r.ID)
	if err != nil {
		return response{}, err
	}

	member.Name = r.Name
	if r.Verified {
		if err = member.MakeVerified(); err != nil {
			return response{}, err
		}
	} else {
		member.MakeUnverified()
	}

	if err = s.repository.Save(&member); err != nil {
		return response{}, err
	}
	return toResponse(member), nil
}

func (s *Service) delete(r findRequest) error {
	return s.repository.Delete(r.ID)
}

func (s *Service) list() ([]response, error) {
	members, err := s.repository.List()
	if err != nil {
		return []response{}, err
	}
	return toListResponse(members), nil
}
