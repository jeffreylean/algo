package algo

// Leetcode problem 649.

//In the world of Dota2, there are two parties: the Radiant and the Dire.

//The Dota2 senate consists of senators coming from two parties. Now the Senate wants to decide on a change in the Dota2 game. The voting for this change is a round-based procedure. In each round, each senator can exercise one of the two rights:
//
//Ban one senator's right: A senator can make another senator lose all his rights in this and all the following rounds.
//Announce the victory: If this senator found the senators who still have rights to vote are all from the same party, he can announce the victory and decide on the change in the game.
//Given a string senate representing each senator's party belonging. The character 'R' and 'D' represent the Radiant party and the Dire party.
//Then if there are n senators, the size of the given string will be n.

//The round-based procedure starts from the first senator to the last senator in the given order. This procedure will last until the end of voting.
//All the senators who have lost their rights will be skipped during the procedure.

//Suppose every senator is smart enough and will play the best strategy for his own party. Predict which party will finally announce the victory and change the Dota2 game.
//The output should be "Radiant" or "Dire".

func predictPartyVictory(senate string) string {
	type senator struct {
		totalCount  int
		bannedCount int
	}

	hm := map[byte]*senator{'R': new(senator), 'D': new(senator)}
	nextRound := make([]byte, 0)
	queue := []byte(senate)

	// get the count of each parties
	for _, each := range senate {
		hm[byte(each)].totalCount++
	}

	for hm['R'].totalCount > 0 && hm['D'].totalCount > 0 {
		for _, each := range queue {
			switch byte(each) {
			case 'R':
				if hm['R'].bannedCount == 0 {
					if hm['D'].bannedCount >= hm['R'].totalCount {
						return "Radiant"
					}
					hm['D'].bannedCount++
					hm['D'].totalCount--
					nextRound = append(nextRound, 'R')
				} else {
					hm['R'].bannedCount--
				}
			case 'D':
				if hm['D'].bannedCount == 0 {
					if hm['R'].bannedCount >= hm['D'].totalCount {
						return "Dire"
					}
					hm['R'].bannedCount++
					hm['R'].totalCount--
					nextRound = append(nextRound, 'D')
				} else {
					hm['D'].bannedCount--
				}
			}
		}
		queue, nextRound = nextRound, queue[:0]
	}
	if hm['R'].totalCount > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}
