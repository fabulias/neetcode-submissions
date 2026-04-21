func numUniqueEmails(emails []string) int {
    unique := make(map[string]bool)

    for _, e := range emails {
        email := ""
        ix := 0
        for e[ix] != '@' && e[ix] != '+' {
            if e[ix] != '.' {
                email += string(e[ix])
            }
            ix++
        }
        for e[ix] != '@' {
            ix++
        }
        domain := e[ix+1:]
        
        key := email + "@" + domain
        unique[key] = true
    }

    return len(unique)
}