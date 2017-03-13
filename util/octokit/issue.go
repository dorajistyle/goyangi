// DEPRECATED BEACAUSE OF code.google.com/p/go-netrc/netrc
package octokit

// import (
// 	"github.com/octokit/go-octokit/octokit"
// 	"net/http"
//   "time"
//   "bytes"
// //   "fmt"
// )

// var (
//   client *octokit.Client
//   owner string
//   repo string
// )

// func setup(gitHubAPIURL, userAgent, accessToken string) *octokit.Client {
// 	httpClient := http.Client{
// 		Transport: http.DefaultTransport,
// 	}
// 	// octokit client configured
// 	client := octokit.NewClientWith(
// 		gitHubAPIURL,
// 		userAgent,
// 		octokit.TokenAuth{AccessToken: accessToken},
// 		&httpClient,
// 	)
// 	return client
// }

// func NewOctokitClient(gitHubAPIURL, userAgent, accessToken, octokitOwner, octokitRepo string) {
//   owner = octokitOwner
//   repo = octokitRepo
//   client = setup(gitHubAPIURL, userAgent, accessToken)
// }

// func searchIssue(client *octokit.Client, query string, level string) (*octokit.IssueSearchResults, *octokit.Result) {
//     queryBuffer := new(bytes.Buffer)
//     queryBuffer.WriteString(query)
//     queryBuffer.WriteString("+label:bug")
//     if len(level) > 0 {
//       queryBuffer.WriteString(" +label:")
//       queryBuffer.WriteString(level)
//     }
//     queryBuffer.WriteString(" +state:open")
//     queryBuffer.WriteString(" +repo:")
//     queryBuffer.WriteString(owner)
//     queryBuffer.WriteString("/")
//     queryBuffer.WriteString(repo)
//     search := client.Search()
//     params := octokit.M{"query": queryBuffer.String()}
//     searchResults, result := search.Issues(nil, params)
//     return searchResults, result
// }

// func createIssue(client *octokit.Client, title string, body string, level string) error{
//   labels := []string{"goyangi-api","bug"}
//   if len(level) > 0 {
//       labels = append(labels, level)
//   }
// 	params := octokit.IssueParams{
// 		Title: title,
// 		Body:  body,
// 	  Labels: labels}
//   time.Now()
// 	_, result := client.Issues().Create(nil, octokit.M{"owner": owner,
// 		"repo": repo}, params)
//   if result.HasError() {
//     return result.Err
//   }
//   return nil
// }

// func updateIssue(client *octokit.Client, number int, level string) error{
//   labels := []string{"goyangi-api", "bug", "duplicate"}
//   if len(level) > 0 {
//       labels = append(labels, level)
//   }
// 	var params octokit.IssueParams
// 	params = octokit.IssueParams{
// 	          Labels: labels}
//   issueInfo := octokit.M{"owner": owner,
// 		"repo": repo, "number": number}
// 	pastIssue, pastIssueResult := client.Issues().One(nil, issueInfo)
// 	if !pastIssueResult.HasError() {
//     now := time.Now()
//     timestamp := now.Format(time.RFC850)
//     bodyBuffer := new(bytes.Buffer)
//     bodyBuffer.WriteString(pastIssue.Body)
//     bodyBuffer.WriteString("\n")
//     bodyBuffer.WriteString(timestamp)
//     body := bodyBuffer.String()
// 		params = octokit.IssueParams{
//  	          Body: body,
// 	          Labels: labels}
// 	}

// 	_, result := client.Issues().Update(nil, issueInfo, params)
//   if result.HasError() {
//     return result.Err
//   }

//   return nil
// }

// func CreateOrUpdateIssue(title string, body string, level string) error {
//     searchResults, result := searchIssue(client, title, level)
//     var err error
//     if !result.HasError() {
//     //   fmt.Printf("searchResults count : %d\n", searchResults.TotalCount)
//        if(searchResults.TotalCount > 0) {
//            number := searchResults.Items[0].Number
//            err = updateIssue(client, number, level)
//        } else {
// 	        err = createIssue(client, title, body, level)
//        }
//     } else {
//     	err = createIssue(client, title, body, level)
//     }
//     return err
// }
