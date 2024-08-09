package models

type TbMahasiswa struct {
	Id  int64 `gorm:"primaryKey" json:"id"`
	Nama string `gorm:"type:varchar(100)" json:"nama" validate:"required"`
	Alamat string `gorm:"type:text" json:"alamat" validate:"required"`
}

func (TbMahasiswa) TableName() string {
	return "tb_mahasiswa"
}

