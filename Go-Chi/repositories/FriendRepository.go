package repositories

import (
	"Go-Chi/interfaces"
	"Go-Chi/models"
	"database/sql"
	"fmt"
)

type Repository struct {
	Db *sql.DB
}

func FriendRepository(db *sql.DB) interfaces.IRepository {
	return &Repository {
		Db: db,
	}
}

func (s *Repository) RCheckNonAddFriend(friends models.Friends) bool {
	connect, err := s.Db.Query("select `user_id` from `connection` where `user_id` = (select `id` from `user` where `email`=?) AND `connect_id` = (select `id` from `user` where `email`=?)", friends.Friends[0], friends.Friends[1])
	catch(err)

	for connect.Next(){
		var userId sql.NullInt64
		err = connect.Scan(&userId)
		catch(err)
		fmt.Println(userId)
		if userId.Valid {
			return false
		}
	}

	defer connect.Close()
	return true
}

func (s *Repository) RAddFriend(friends models.Friends) error{
	addFriend, err := s.Db.Prepare("INSERT `connection` SET user_id=(select `id` from `user` where `email`=?), connect_id=(select `id` from `user` where `email`=?)")
	catch(err)
	_, err = addFriend.Exec(friends.Friends[0], friends.Friends[1])
	catch(err)
	_, err = addFriend.Exec(friends.Friends[1], friends.Friends[0])
	catch(err)
	defer addFriend.Close()
	return err
}

func (s *Repository) RFindFriendsOfUser(m models.Email) []string {
	emailFriends, err := s.Db.Query("select `email` from `user` where `id` in (select `connect_id` from `connection` where `user_id` = (select `id` from `user` where `email`=?))", m.Email)
	var email []string
	for emailFriends.Next(){
		var e string
		err = emailFriends.Scan(&e)
		catch(err)
		email = append(email, e)
	}
	return email
}

func (s *Repository) RFindCommonFriends(friends models.Friends)[]string{
	commonFriends, err := s.Db.Query("SELECT `email` from `user` WHERE `id` IN (SELECT `user_id` from `connection` JOIN (SELECT `id` FROM `user` where `email` in ( ?, ?)) t ON `connect_id` = `id` group by `user_id` having count(`user_id`) > 1)", friends.Friends[0], friends.Friends[1])
	catch(err)
	var email []string
	for commonFriends.Next(){
		var e string
		commonFriends.Scan(&e)
		email = append(email, e)
	}
	return email
}

func (s *Repository) RCheckNonFollow(subscribe models.Request) bool {
	follow, err := s.Db.Query("select `user_id` from `follow` where `user_id` = (select `id` from `user` where `email`=?) AND `follow_id` = (select `id` from `user` where `email`=?)", subscribe.Requestor, subscribe.Target)
	catch(err)

	for follow.Next(){
		var userId sql.NullInt64
		err = follow.Scan(&userId)
		catch(err)
		fmt.Println(userId)
		if userId.Valid {
			return false
		}
	}

	defer follow.Close()
	return true
}

func (s *Repository) RFollowFriend(subscribe models.Request) error {
	followUser, err := s.Db.Prepare("INSERT `follow` SET `user_id`=(select `id` from `user` where `email`=?), follow_id=(select `id` from `user` where `email`=?)")
	catch(err)
	_, err = followUser.Exec(subscribe.Requestor, subscribe.Target)
	catch(err)
	defer followUser.Close()
	return err
}

func (s *Repository) RCheckNonBlock(subscribe models.Request) bool {
	block, err := s.Db.Query("select `user_id` from `block` where `user_id` = (select `id` from `user` where `email`=?) AND `block_id` = (select `id` from `user` where `email`=?)", subscribe.Requestor, subscribe.Target)
	catch(err)

	for block.Next(){
		var userId sql.NullInt64
		err = block.Scan(&userId)
		catch(err)
		fmt.Println(userId)
		if userId.Valid {
			return false
		}
	}

	defer block.Close()
	return true
}

func (s *Repository) RBlockFriend(subscribe models.Request) error {
	blockUser, err := s.Db.Prepare("INSERT `block` SET `user_id`=(select `id` from `user` where `email`=?), block_id=(select `id` from `user` where `email`=?)")
	catch(err)
	_, err = blockUser.Exec(subscribe.Requestor, subscribe.Target)
	catch(err)
	defer blockUser.Close()
	return err
}

func (s *Repository) RNonBlockByEmail(sender models.Sender) []string {
	nonBlockId, err := s.Db.Query("SELECT `email` FROM `user` WHERE `id` NOT IN (SELECT `block_id` from `block` join( SELECT `id` FROM `user` where `email` = ?) `u` ON `user_id` = `u`.`id`)", sender.Sender)
	catch(err)
	var emails []string
	for nonBlockId.Next() {
		var email string
		err = nonBlockId.Scan(&email)
		catch(err)
		emails = append(emails, email)
	}
	return emails
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}