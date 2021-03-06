package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	mode string
	idst string
	ls   string
)

var rootCmd = &cobra.Command{
	Use: "app",
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: func(cmd *cobra.Command, args []string) {

		if mode == "" {
			mode = MODE_BASIC
		}

		if idst == "" {
			fmt.Println("university ID cannot be empty!")
			os.Exit(1)
		}

		if ls == "" {
			ls = abc
		}

		if mode == MODE_BASIC {

			col := []string{"ID", "NAMA", "NIP", "PERGURUAN TINGGI", "JENJANG", "PROGRAM STUDI"}
			if err := addcol("dosen.csv", col); err != nil {
				log.Fatal(err)
			}

			for _, v := range ls {
				fmt.Println(fmt.Sprintf("Lecture name contain letter %c :", v))
				err := reqDosen(fmt.Sprintf("%c", v), idst, func(d Dosen) {

					fmt.Println("Lecture name : ", d.Nama, " Added to dosen.csv")
					col := []string{d.ID, d.Nama, d.Nip, d.Pt, d.Jenjang, d.Prodi}
					exist, err := checkcol("dosen.csv", col)
					if err != nil {
						log.Fatal(err)
					}
					if !exist {
						if err := addcol("dosen.csv", col); err != nil {
							log.Fatal(err)
						}
					}
				})
				if err != nil {
					log.Fatal(err)
				}
			}

		} else if mode == MODE_ADVANCE {

			if err := os.MkdirAll("dosen", os.ModePerm); err != nil {
				log.Fatal(err)
			}

			for _, v := range ls {
				fmt.Println(fmt.Sprintf("Lecture name contain letter %c :", v))
				err := reqDosen(fmt.Sprintf("%c", v), idst, func(d Dosen) {

					fmt.Println("Lecture name : ", d.Nama, " Added to ", fmt.Sprintf("dosen/%s/DataUmum.csv", d.Nama))
					if err := reqDosenDetail(d.ID, func(detail DosenDetail) {

						if err := os.RemoveAll(fmt.Sprintf("dosen/%s", d.Nama)); err != nil {
							log.Fatal(err)
						}

						if err := os.MkdirAll(fmt.Sprintf("dosen/%s", d.Nama), os.ModePerm); err != nil {
							log.Fatal(err)
						}

						coldmHeader := []string{"ID MATAKULIAH", "NOMOR KELAS", "KODE KELAS", "NAMA MATAKULIAH", "NAMA PERGURUAN TINGGI", "LINK PERGURUAN TINGGI"}
						if err := addcol(fmt.Sprintf("dosen/%s/DataMengajar.csv", d.Nama), coldmHeader); err != nil {
							log.Fatal(err)
						}

						for _, dm := range detail.Datamengajar {
							col := []string{
								dm.IDSmt,
								dm.NmKls,
								dm.KodeMk,
								dm.NmMk,
								dm.Namapt,
								dm.Linkpt,
							}
							if err := addcol(fmt.Sprintf("dosen/%s/DataMengajar.csv", d.Nama), col); err != nil {
								log.Fatal(err)
							}
						}

						coldpHeader := []string{"TAHUN LULUS", "NAMA SP FORMAL", "NAMA JENJANG", "SINGKAT GELAR"}
						if err := addcol(fmt.Sprintf("dosen/%s/DataPendidikan.csv", d.Nama), coldpHeader); err != nil {
							log.Fatal(err)
						}

						for _, dp := range detail.Datapendidikan {
							col := []string{
								fmt.Sprint(dp.ThnLulus),
								dp.NmSpFormal,
								dp.Namajenjang,
								dp.SingkatGelar,
							}
							if err := addcol(fmt.Sprintf("dosen/%s/DataPendidikan.csv", d.Nama), col); err != nil {
								log.Fatal(err)
							}
						}

						coldumumHeader := []string{
							"ID SDM",
							"NAMA SDM",
							"JENIS KELAMIN",
							"TEMPAT LAHIT",
							"NAMA PERGURUAN TINGGI",
							"LINK PERGURUAN TINGGI",
							"LINK PRODI",
							"NAMA PRODI",
							"STATUS KEAKTIFAN",
							"PENDIDIKAN TINGGI",
							"FUNGSIONAL",
							"IKATAN KERJA",
						}
						if err := addcol(fmt.Sprintf("dosen/%s/DataUmum.csv", d.Nama), coldumumHeader); err != nil {
							log.Fatal(err)
						}

						col := []string{
							detail.Dataumum.IDSdm,
							detail.Dataumum.NmSdm,
							detail.Dataumum.Jk,
							detail.Dataumum.TmptLahir,
							detail.Dataumum.Namapt,
							detail.Dataumum.Linkpt,
							detail.Dataumum.Linkprodi,
							detail.Dataumum.Namaprodi,
							detail.Dataumum.Statuskeaktifan,
							detail.Dataumum.PendTinggi,
							detail.Dataumum.Fungsional,
							detail.Dataumum.Ikatankerja,
						}
						if err := addcol(fmt.Sprintf("dosen/%s/DataUmum.csv", d.Nama), col); err != nil {
							log.Fatal(err)
						}

						colBasicHeader := []string{"ID", "NAMA", "NIP", "PERGURUAN TINGGI", "JENJANG", "PROGRAM STUDI"}
						if err := addcol(fmt.Sprintf("dosen/%s/DataBasic.csv", d.Nama), colBasicHeader); err != nil {
							log.Fatal(err)
						}

						colBasic := []string{d.ID, d.Nama, d.Nip, d.Pt, d.Jenjang, d.Prodi}
						if err := addcol(fmt.Sprintf("dosen/%s/DataBasic.csv", d.Nama), colBasic); err != nil {
							log.Fatal(err)
						}

					}); err != nil {
						log.Fatal(err)
					}

				})
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&mode, "mode", "", "mode access (can not be empty)")
	rootCmd.PersistentFlags().StringVar(&idst, "univ", "", "university id (must not be empty)")
	rootCmd.PersistentFlags().StringVar(&ls, "ls", "", "letter search (can be empty)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
