package player

import "errors"

type Stats struct {
  Name string
  Minutes float32
  Points,Assists,TurnOver,Rebounds int8
}

func HadAGoodGame(_stats Stats)(bool, error){
  if _stats.Assists < 0 || _stats.Minutes < 0 ||_stats.Points < 0 ||_stats.TurnOver <0 || _stats.Rebounds < 0{
    return false, errors.New("Stats can't be negative")
  }
  if _stats.Name == ""{
   return false, errors.New("name can't be empty")
  }
  // Define criteria for a good game (example criteria, can be adjusted)
  goodGamePoints := 20
  goodGameAssists := 5
  goodGameRebounds := 10

  // Check if the player had a good game based on the criteria
  if _stats.Points >= int8(goodGamePoints) || _stats.Assists >= int8(goodGameAssists) || _stats.Rebounds >= int8(goodGameRebounds) {
      return true, nil
  }
  // If criteria are not met, return false
  return false, nil
}