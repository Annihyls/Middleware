from marshmallow import Schema, fields, validates_schema, ValidationError


# Schéma rating de sortie (renvoyé au front)
class RatingSchema(Schema):
    note = fields.Int(description="Note")
    description = fields.String(description="Description")

    @staticmethod
    def is_empty(obj):
        return (not obj.get("note"))


class BaseRatingSchema(Schema):
    note = fields.Int(description="Note")
    description = fields.String(description="Description")


# Schéma rating de modification (note, description)
class RatingUpdateSchema(BaseRatingSchema):
    # permet de définir dans quelles conditions le schéma est validé ou non
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not ("note" in data and data["note"] is not None):
            raise ValidationError("at least one of ['note'] must be specified")


class RatingCreateSchema(BaseRatingSchema):
    # permet de définir dans quelles conditions le schéma est validé ou non
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("note" in data and data["note"] is not None) or
                ("id_song" in data and data["id_song"] != "") or
                ("id_user" in data and data["id_user"] != "")):
            raise ValidationError("at least one of ['note','id_user','id_song'] must be specified")

