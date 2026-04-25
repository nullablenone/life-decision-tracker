package decision

import "gorm.io/gorm"

func SeedDecisions(db *gorm.DB) error {
	decisions := []DecisionCategory{
		{Name: "Brilliant", Value: 5, Description: "Keputusan langka yang berdampak sangat besar dan positif di luar dugaan.", Color: "#1baaa0"},
		{Name: "Great", Value: 4, Description: "Langkah yang sangat solid dan memberikan keuntungan signifikan.", Color: "#5c8bb0"},
		{Name: "Best", Value: 3, Description: "Pilihan terbaik yang paling optimal dalam situasi tersebut.", Color: "#81b64c"},
		{Name: "Excellent", Value: 2, Description: "Keputusan yang sangat bagus dan mendukung tujuan utama.", Color: "#96bc4b"},
		{Name: "Good", Value: 1, Description: "Langkah yang benar, meskipun mungkin ada opsi yang sedikit lebih baik.", Color: "#96af8b"},
		{Name: "Book", Value: 0, Description: "Keputusan standar, rutinitas, atau kewajiban yang memang sudah seharusnya dilakukan.", Color: "#a88865"},
		{Name: "Inaccuracy", Value: -1, Description: "Keputusan yang kurang optimal, sedikit membuang waktu atau sumber daya.", Color: "#f0c15c"},
		{Name: "Mistake", Value: -2, Description: "Pilihan yang salah dan memberikan dampak negatif yang nyata.", Color: "#e6912c"},
		{Name: "Miss", Value: -3, Description: "Melewatkan kesempatan penting atau mengabaikan prioritas utama.", Color: "#ff7769"},
		{Name: "Blunder", Value: -5, Description: "Kesalahan fatal yang merusak progres atau melanggar prinsip dasar.", Color: "#fa412d"},
	}

	for _, d := range decisions {
		err := db.FirstOrCreate(&d, DecisionCategory{Name: d.Name, Description: d.Description, Value: d.Value, Color: d.Color}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
