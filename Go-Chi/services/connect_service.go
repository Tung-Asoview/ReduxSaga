package services

import (
	"Go-Chi/Model"
	"Go-Chi/driver"
	"Go-Chi/repository"
	"encoding/json"
	"net/http"
)

func AddFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	connectionRepo := repository.FriendRepo(db)
	var friends Model.Friends
	json.NewDecoder(r.Body).Decode(&friends)
	var block = Model.Request{Requestor:friends.Friends[0], Target:friends.Friends[1]}
	var beBlock = Model.Request{Requestor:friends.Friends[1], Target:friends.Friends[0]}

	var checkNonAddFriend = connectionRepo.CheckNonAddFriend(friends)
	var checkNonBlock = connectionRepo.CheckNonBlock(block)
	var checkNonBeBlock = connectionRepo.CheckNonBlock(beBlock)

	if checkNonAddFriend && checkNonBlock && checkNonBeBlock{
		connectionRepo.AddFriend(friends)
		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": "failed add friend"})
	}
}

func FindFriendsOfUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	connectionRepo := repository.FriendRepo(db)
	var mail Model.Email
	json.NewDecoder(r.Body).Decode(&mail)
	var status = Model.Friends{Success: true,Friends: connectionRepo.FindFriendsOfUser(mail), Count: len(connectionRepo.FindFriendsOfUser(mail))}
	if len(status.Friends) >0 {
		respondwithJSON(w, http.StatusOK, status)
	} else {
		respondwithJSON(w, http.StatusNotFound, map[string]string{"message": "Not Found Friends Of User"})
	}
}

func FindCommonFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	connectionRepo := repository.FriendRepo(db)
	var friends Model.Friends
	json.NewDecoder(r.Body).Decode(&friends)
	var status = Model.Friends{Success: true, Friends: connectionRepo.FindCommonFriends(friends), Count: len(connectionRepo.FindCommonFriends(friends))}
	if len(status.Friends) >0 {
		respondwithJSON(w, http.StatusOK, status)
	} else {
		respondwithJSON(w, http.StatusNotFound, map[string]string{"message": "Not Found Common Friends"})
	}
}

func FollowFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	connectionRepo := repository.FriendRepo(db)
	var subscribe Model.Request
	json.NewDecoder(r.Body).Decode(&subscribe)
	//var block = Model.Request{Requestor:subscribe.Requestor, Target:subscribe.Target}
	//var beBlock = Model.Request{Requestor:subscribe.Target, Target:subscribe.Requestor}

	var checkNonFollow = connectionRepo.CheckNonFollow(subscribe)
	//var checkNonBlock = connectionRepo.CheckNonBlock(block)
	//var checkNonBeBlock = connectionRepo.CheckNonBlock(beBlock)

	if checkNonFollow{
		connectionRepo.Follow(subscribe)
		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": "failed follow"})
	}
}

func BlockFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	connectionRepo := repository.FriendRepo(db)
	var block Model.Request
	json.NewDecoder(r.Body).Decode(&block)

	//var friends = Model.Friends{Friends:[]string{block.Requestor, block.Target}}
	//var beFriend = Model.Friends{Friends:[]string{block.Target, block.Requestor}}
	//
	//var checkNonAddFriend = connectionRepo.CheckNonAddFriend(friends)
	//if !checkNonAddFriend {
	//	connectionRepo.DeleteFriends(friends)
	//	connectionRepo.DeleteFriends(beFriend)
	//}
	//
	//var checkNonFollow = connectionRepo.CheckNonFollow(block)
	//if !checkNonFollow {
	//	connectionRepo.DeleteFollow(block)
	//}
	//
	//var beFollow = Model.Request{Requestor:block.Target, Target:block.Requestor}
	//var checkNonBeFollow = connectionRepo.CheckNonFollow(beFollow)
	//if !checkNonBeFollow {
	//	connectionRepo.DeleteFollow(beFollow)
	//}

	var checkNonBlock = connectionRepo.CheckNonBlock(block)
	if checkNonBlock {
		connectionRepo.Block(block)
		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"failed": "You were blocked this account !!!"})
	}
}

func receiveUpdatesFromEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	connectionRepo := repository.FriendRepo(db)
	var recipients Model.Recipients
	json.NewDecoder(r.Body).Decode(&recipients)
	var emails = connectionRepo.NonBlockByEmail(recipients)


	if checkNonFollow && checkNonBlock && checkNonBeBlock{
		connectionRepo.Follow(subscribe)
		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": "failed follow"})
	}
}