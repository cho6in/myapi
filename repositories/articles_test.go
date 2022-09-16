package repositories_test

import (
  "database/sql"
  "testing"
  "fmt"

  "github.com/cho6in/myapi/models"
  "github.com/cho6in/myapi/repositories"
  
  _ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
  dbUser := "docker"
  dbPassword := "docker"
  dbDatabase := "sampledb"
  dbConn := fmt.Sprintf("%s:%s@tpc(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

  db ,err := sql.Open("mysql", dbConn)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  expected := models.Article{
    ID:       1,
    Title:    "firstPost",
    Contents: "This is my first blog",
    UserName: "saki",
    NiceNum:  2,
  }

  got, err := repositories.SelectArticleDetail(db, expected.ID)
  if err != nil {
    t.Fatal(err)
  }

  if got.ID != expected.ID {
    t.Errorf("ID: get %d but want %d\n", got.ID, expected.ID)
  }

  if got.Title != expected.Title {
    t.Errorf("Title get %s but want %s\n", got.Title, expected.Title)
  }

  if got.Contents != expected.Contents {
    t.Errorf("Contents get %s but want %s\n", got.Contents, expected.Contents)
  }

  if got.UserName != expected.UserName {
    t.Errorf("UserName get %s but want %s\n", got.UserName, expected.UserName)
  }

  if got.NiceNum != expected.NiceNum {
    t.Errorf("NiceNum get %d but want %d\n", got.NiceNum, expected.NiceNum)
  }
}

