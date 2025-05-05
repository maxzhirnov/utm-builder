package utm

type UTMParams struct {
    URL        string
    Source     string
    Medium     string
    Term       string
    Content    string
    Campaign   string
}

func (u *UTMParams) BuildLink() string {
    // Собирает ссылку с UTM-метками
    params := ""
    if u.Source != "" {
        params += "&utm_source=" + u.Source
    }
    if u.Medium != "" {
        params += "&utm_medium=" + u.Medium
    }
    if u.Term != "" {
        params += "&utm_term=" + u.Term
    }
    if u.Content != "" {
        params += "&utm_content=" + u.Content
    }
    if u.Campaign != "" {
        params += "&utm_campaign=" + u.Campaign
    }
    if params != "" {
        params = "?" + params[1:] // убираем первый &
    }
    return u.URL + params
}
