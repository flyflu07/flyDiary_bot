package query

var FindMessageQuery = "SELECT d.id, d.text, d.recording_time + (interval '1 hour' * p.user_timezone) FROM diary d JOIN profile p ON d.id = p.id WHERE p.id = @id AND d.recording_time >= @date  AND d.recording_time < @nextday"

var SaveDiaryQuery = "insert into diary values(@id, @message, @user_time, @date);"

var GetUniqueYearsQuery = "select distinct extract(year from recording_time ) from diary where id=@id order by 1;"

var GetUniqueMonthsQuery = "select distinct extract(month from recording_time ) from diary where id=@id and extract(year from recording_time )=@year order by 1;"

var GetUniqueDaysQuery = "select distinct extract(day from recording_time ) from diary where id=@id and extract(year from recording_time )=@year and extract(month from recording_time)=@month order by 1;"
