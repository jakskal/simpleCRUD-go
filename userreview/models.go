package userreview

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jakskal/simpleCRUD-go/config"
	"github.com/julienschmidt/httprouter"
)

type review struct {
	ID        int     `json:"id"`
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	UserID    int     `json:"user_id"`
	Rating    float64 `json:"rating"`
	Review    string  `json:"review"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func allReview(c chan []review) ([]review, error) {
	rows, err := config.DB.Query("SELECT * FROM user_reviews")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	response := make([]review, 0)
	for rows.Next() {
		Review := review{}
		err = rows.Scan(&Review.ID, &Review.OrderID, &Review.ProductID, &Review.UserID, &Review.Rating, &Review.Review, &Review.CreatedAt, &Review.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		response = append(response, Review)
	}
	c <- response
	return response, nil
}

func createReview(r *http.Request) (review, error) {
	newReview := review{}

	json.NewDecoder(r.Body).Decode(&newReview)
	if newReview.Rating < 1 || newReview.Rating > 5 {
		err := errors.New("rating must be between 1 and 5")
		log.Panicln(err)
		return newReview, err
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return newReview, err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO user_reviews(id, order_id, product_id, user_id, rating, review ) values(?,?,?,?,?,?)")
	if err != nil {
		return newReview, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(newReview.ID, newReview.OrderID, newReview.ProductID, newReview.UserID, newReview.Rating, newReview.Review)
	if err != nil {
		return newReview, err
	}

	err = tx.Commit()
	if err != nil {
		return newReview, err
	}

	resultID, err := result.LastInsertId()
	if err != nil {
		return newReview, err
	}

	newReview.ID = int(resultID)

	return newReview, nil
}

func updateReview(r *http.Request) (review, error) {
	updatedReview := review{}

	json.NewDecoder(r.Body).Decode(&updatedReview)
	if updatedReview.Rating < 1 || updatedReview.Rating > 5 {
		err := errors.New("rating must be between 1 and 5")
		log.Panicln(err)
		return updatedReview, err
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return updatedReview, err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("UPDATE user_reviews SET order_id=?, product_id=?, user_id=?, rating=?, review=?  WHERE id=?")
	if err != nil {
		return updatedReview, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(updatedReview.OrderID, updatedReview.ProductID, updatedReview.UserID, updatedReview.Rating, updatedReview.Review, updatedReview.ID)
	if err != nil {
		return updatedReview, err
	}

	affected, _ := result.RowsAffected()
	fmt.Println(affected)

	err = tx.Commit()
	if err != nil {
		return updatedReview, err
	}

	return updatedReview, err
}

func deleteReview(p httprouter.Params) (map[string]int, error) {
	response := make(map[string]int, 0)
	tx, err := config.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	id := p.ByName("id")
	if err != nil {
		return response, err
	}
	result, err := tx.Exec("DELETE FROM user_reviews where id=?", id)
	if err != nil {
		log.Fatal(err)
		return response, err
	}

	err = tx.Commit()
	if err != nil {
		return response, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return response, err
	}

	ids, err := strconv.Atoi(id)

	response["affected row"] = int(affected)
	response["deleted review id"] = ids
	return response, nil
}

func findReview(p httprouter.Params) (review, error) {
	Review := review{}

	id := p.ByName("id")
	err := config.DB.QueryRow("SELECT * FROM user_reviews WHERE id=?", id).Scan(&Review.ID, &Review.OrderID, &Review.ProductID, &Review.UserID, &Review.Rating, &Review.Review, &Review.CreatedAt, &Review.UpdatedAt)
	if err != nil {
		return Review, err
	}

	return Review, err
}
