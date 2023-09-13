package main

type Movie struct {
	ID       string   `json:"id"`
	ISBN     string   `json:"isbn"`
	Title    string   `json:"title"`
	Director Director `json:"director"`
}

type Director struct {
	Name string `json:"name"`
}

func FakeDB()  {
	Movies = append(Movies,
		Movie{
			ID:    "1",
			ISBN:  "13: 9780762459469",
			Title: "Bal Hanuman",
			Director: Director{
				Name: "Sandeep chawla",
			},
		},
		Movie{
			ID:    "2",
			ISBN:  "13: 9780762459749",
			Title: "Bal Ganesh",
			Director: Director{
				Name: "Ayush patil",
			},
		},
		Movie{
			ID:    "3",
			ISBN:  "13: 978076245784",
			Title: "chota bheem",
			Director: Director{
				Name: "shivam verma",
			},
		},
	)

}