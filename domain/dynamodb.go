package domain

type Data struct {
	Name string `dynamo:"name,hash"`
	Text string `dyname:"text"`
}
