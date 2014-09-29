package filepath

import "path/filepath"
import "net/url"

// Base parses the base name from a url.URL or string (for fs)
func Base(o interface{}) string {
	var name string

	switch v := o.(type) {
	case string:
		return filepath.Base(v)
	case *url.URL:
		name = filepath.Base(v.Path)
		if "." == name || "/" == name {
			name = v.Host
		}
	}

	return name
}
