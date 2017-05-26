package main

import (
	"net/http"
	"time"
	"encoding/json"
	"log"
)

type gameData struct {
	Subject string `json:"subject"`
	Copyright string `json:"copyright"`
	Data struct {
		Game struct {
			HomeTeamRuns string `json:"home_team_runs"`
			GameType string `json:"game_type"`
			DoubleHeaderSw string `json:"double_header_sw"`
			Location string `json:"location"`
			HomeRecapLink string `json:"home_recap_link"`
			AwayTime string `json:"away_time"`
			AwayPreviewLink string `json:"away_preview_link"`
			IsNoHitter string `json:"is_no_hitter"`
			Time string `json:"time"`
			TopInning string `json:"top_inning"`
			HomeTime string `json:"home_time"`
			HomeTeamName string `json:"home_team_name"`
			Ind string `json:"ind"`
			OriginalDate string `json:"original_date"`
			HomeTeamCity string `json:"home_team_city"`
			VenueID string `json:"venue_id"`
			GamedaySw string `json:"gameday_sw"`
			AwayWin string `json:"away_win"`
			HomeGamesBackWildcard string `json:"home_games_back_wildcard"`
			SavePitcher struct {
				Last string `json:"last"`
				NameDisplayRoster string `json:"name_display_roster"`
				Era string `json:"era"`
				ID string `json:"id"`
				FirstName string `json:"first_name"`
				SLosses string `json:"s_losses"`
				SEra string `json:"s_era"`
				Saves string `json:"saves"`
				LastName string `json:"last_name"`
				Losses string `json:"losses"`
				First string `json:"first"`
				SWins string `json:"s_wins"`
				Wins string `json:"wins"`
			} `json:"save_pitcher"`
			AwayTeamID string `json:"away_team_id"`
			TzHmLgGen string `json:"tz_hm_lg_gen"`
			Status string `json:"status"`
			HomeLoss string `json:"home_loss"`
			HomeGamesBack string `json:"home_games_back"`
			HomeCode string `json:"home_code"`
			AwaySportCode string `json:"away_sport_code"`
			HomeWin string `json:"home_win"`
			TimeHmLg string `json:"time_hm_lg"`
			AwayNameAbbrev string `json:"away_name_abbrev"`
			League string `json:"league"`
			TimeZoneAwLg string `json:"time_zone_aw_lg"`
			AwayGamesBack string `json:"away_games_back"`
			HomeFileCode string `json:"home_file_code"`
			GameDataDirectory string `json:"game_data_directory"`
			TimeZone string `json:"time_zone"`
			AwayLeagueID string `json:"away_league_id"`
			Preview string `json:"preview"`
			IsPerfectGame string `json:"is_perfect_game"`
			AwayRecapLink string `json:"away_recap_link"`
			HomeTeamID string `json:"home_team_id"`
			TimeAwLg string `json:"time_aw_lg"`
			AwayTeamCity string `json:"away_team_city"`
			Day string `json:"day"`
			TbdFlag string `json:"tbd_flag"`
			TzAwLgGen string `json:"tz_aw_lg_gen"`
			AwayCode string `json:"away_code"`
			WinningPitcher struct {
				ID string `json:"id"`
				SLosses string `json:"s_losses"`
				FirstName string `json:"first_name"`
				SEra string `json:"s_era"`
				Last string `json:"last"`
				LastName string `json:"last_name"`
				Losses string `json:"losses"`
				Era string `json:"era"`
				NameDisplayRoster string `json:"name_display_roster"`
				First string `json:"first"`
				SWins string `json:"s_wins"`
				Wins string `json:"wins"`
			} `json:"winning_pitcher"`
			GameMedia struct {
				Media struct {
					Free string `json:"free"`
					Title string `json:"title"`
					Thumbnail string `json:"thumbnail"`
					MediaState string `json:"media_state"`
					Start string `json:"start"`
					HasMlbtv string `json:"has_mlbtv"`
					CalendarEventID string `json:"calendar_event_id"`
					Enhanced string `json:"enhanced"`
					Type string `json:"type"`
				} `json:"media"`
			} `json:"game_media"`
			GameNbr string `json:"game_nbr"`
			TimeDateAwLg string `json:"time_date_aw_lg"`
			AwayGamesBackWildcard string `json:"away_games_back_wildcard"`
			ScheduledInnings string `json:"scheduled_innings"`
			Linescore []struct {
				AwayInningRuns string `json:"away_inning_runs"`
				Inning string `json:"inning"`
				HomeInningRuns string `json:"home_inning_runs"`
			} `json:"linescore"`
			VenueWChanLoc string `json:"venue_w_chan_loc"`
			FirstPitchEt string `json:"first_pitch_et"`
			AwayTeamName string `json:"away_team_name"`
			GamedayLink string `json:"gameday_link"`
			Outs string `json:"outs"`
			TimeDateHmLg string `json:"time_date_hm_lg"`
			ID string `json:"id"`
			HomeNameAbbrev string `json:"home_name_abbrev"`
			WrapupLink string `json:"wrapup_link"`
			TiebreakerSw string `json:"tiebreaker_sw"`
			Balls string `json:"balls"`
			Ampm string `json:"ampm"`
			HomeDivision string `json:"home_division"`
			HomeTimeZone string `json:"home_time_zone"`
			TvStation string `json:"tv_station"`
			AwayTeamErrors string `json:"away_team_errors"`
			AwayTeamRuns string `json:"away_team_runs"`
			Note string `json:"note"`
			AwayTimeZone string `json:"away_time_zone"`
			HmLgAmpm string `json:"hm_lg_ampm"`
			HomeSportCode string `json:"home_sport_code"`
			TimeDate string `json:"time_date"`
			Inning string `json:"inning"`
			HomeTeamHits string `json:"home_team_hits"`
			InningState string `json:"inning_state"`
			HomeAmpm string `json:"home_ampm"`
			GamePk string `json:"game_pk"`
			Venue string `json:"venue"`
			HomeLeagueID string `json:"home_league_id"`
			AwayTeamHits string `json:"away_team_hits"`
			AwayLoss string `json:"away_loss"`
			HomeTeamErrors string `json:"home_team_errors"`
			AwayFileCode string `json:"away_file_code"`
			LosingPitcher struct {
				ID string `json:"id"`
				SLosses string `json:"s_losses"`
				FirstName string `json:"first_name"`
				SEra string `json:"s_era"`
				Last string `json:"last"`
				LastName string `json:"last_name"`
				Losses string `json:"losses"`
				Era string `json:"era"`
				NameDisplayRoster string `json:"name_display_roster"`
				First string `json:"first"`
				SWins string `json:"s_wins"`
				Wins string `json:"wins"`
			} `json:"losing_pitcher"`
			Strikes string `json:"strikes"`
			AwLgAmpm string `json:"aw_lg_ampm"`
			HomePreviewLink string `json:"home_preview_link"`
			TimeZoneHmLg string `json:"time_zone_hm_lg"`
			AwayAmpm string `json:"away_ampm"`
			AwayDivision string `json:"away_division"`
		} `json:"game"`
	} `json:"data"`
}

func main() {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	var url = "http://gd2.mlb.com/components/game/mlb/year_2017/month_05/day_09/gid_2017_05_09_bosmlb_milmlb_1/boxscore.json"
	r, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(gameData)

}

