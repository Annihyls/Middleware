from marshmallow import Schema, fields, validates_schema, ValidationError


# Schéma utilisateur de sortie (renvoyé au front)
class SongSchema(Schema):
    name = fields.String(description="title")
    songname = fields.String(description="artist")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("title") or obj.get("artist") == "")

class BaseSongSchema(Schema):
    title = fields.String(description="title")
    artist = fields.String(description="artist")


# Schéma utilisateur de modification (name, songname, password)
class SongUpdateSchema(BaseSongSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("title" in data and data["title"] != "") or
                ("artist" in data and data["artist"] != "")):
            raise ValidationError("at least one of ['title','artist' must be specified")
    
class SongCreateSchema(BaseSongSchema):
    # permet de définir dans quelles conditions le schéma est validé ou non lorsqu'on le create
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("title" in data and data["title"] != "") or
                ("artist" in data and data["artist"] != "")):
            raise ValidationError("at least one of ['title','artist'] must be specified")

