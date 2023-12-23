from marshmallow import Schema, fields, validates_schema, ValidationError


# Schéma rating de sortie (renvoyé au front)
class RatingSchema(Schema):
    note = fields.String(description="Note")
    description = fields.String(description="Description")

    @staticmethod
    def is_empty(obj):
        return (not obj.get("note") or obj.get("note") == "") and \
               (not obj.get("description") or obj.get("description") == "")


class BaseRatingSchema(Schema):
    note = fields.String(description="Name")
    description = fields.String(description="Description")


# Schéma rating de modification (note, description)
class RatingUpdateSchema(BaseRatingSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("note" in data and data["note"] != "") or
                ("description" in data and data["description"] != "")):
            raise ValidationError("at least one of ['note','description'] must be specified")
