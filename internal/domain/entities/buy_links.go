package entities

type BuyLinks struct {
	links []BuyLink
}

func NewBuyLinks(links []BuyLink) BuyLinks {
	return BuyLinks{links: links}
}

func (b BuyLinks) Links() []BuyLink {
	return b.links
}

func (b BuyLinks) AddLink(link BuyLink) BuyLinks {
	b.links = append(b.links, link)
	return b
}

func (b BuyLinks) RemoveLink(link BuyLink) BuyLinks {
	for i, l := range b.links {
		if l == link {
			b.links = append(b.links[:i], b.links[i+1:]...)
		}
	}
	return b
}

func (b BuyLinks) ActivateLink(link BuyLink) BuyLinks {
	for i, l := range b.links {
		if l == link {
			b.links[i].active = true
		}
	}
	return b
}

func (b BuyLinks) DeactivateLink(link BuyLink) BuyLinks {
	for i, l := range b.links {
		if l == link {
			b.links[i].active = false
		}
	}
	return b
}

func (b BuyLinks) GetActiveLinks() []BuyLink {
	var activeLinks []BuyLink
	for _, link := range b.links {
		if link.active {
			activeLinks = append(activeLinks, link)
		}
	}
	return activeLinks
}

func (b BuyLinks) GetInactiveLinks() []BuyLink {
	var inactiveLinks []BuyLink
	for _, link := range b.links {
		if !link.active {
			inactiveLinks = append(inactiveLinks, link)
		}
	}
	return inactiveLinks
}
