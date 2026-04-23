package decision

import "gorm.io/gorm"

func SeedDecisions(db *gorm.DB) error {
	decisions := []DecisionCategory{
		{Name: "Brilliant", Value: 5, Description: "Keputusan langka yang berdampak sangat besar dan positif di luar dugaan."},
		{Name: "Great", Value: 4, Description: "Langkah yang sangat solid dan memberikan keuntungan signifikan."},
		{Name: "Best", Value: 3, Description: "Pilihan terbaik yang paling optimal dalam situasi tersebut."},
		{Name: "Excellent", Value: 2, Description: "Keputusan yang sangat bagus dan mendukung tujuan utama."},
		{Name: "Good", Value: 1, Description: "Langkah yang benar, meskipun mungkin ada opsi yang sedikit lebih baik."},
		{Name: "Book", Value: 0, Description: "Keputusan standar, rutinitas, atau kewajiban yang memang sudah seharusnya dilakukan."},
		{Name: "Inaccuracy", Value: -1, Description: "Keputusan yang kurang optimal, sedikit membuang waktu atau sumber daya."},
		{Name: "Mistake", Value: -2, Description: "Pilihan yang salah dan memberikan dampak negatif yang nyata."},
		{Name: "Miss", Value: -3, Description: "Melewatkan kesempatan penting atau mengabaikan prioritas utama."},
		{Name: "Blunder", Value: -5, Description: "Kesalahan fatal yang merusak progres atau melanggar prinsip dasar."},
	}

	for _, d := range decisions {
		err := db.FirstOrCreate(&d, DecisionCategory{Name: d.Name}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
