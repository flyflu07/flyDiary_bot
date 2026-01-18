package query

var PasswordCheckQuery = "select exists(select 1 from profile where id=@id and ispassword)"

var CreateProfileQuery = "insert into profile(id, ispassword, passmd5) values (@id, true, @password);"

var GetPasswordQuery = "select passmd5 from profile p  where p.id = @id;"

var SaveTimeZone = "update profile set timezone=@timezone where id=@id"

var GetTimeZoneQuery = "select timezone from profile where id=@id;"
