import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from src.schemas.rating import RatingSchema
from src.models.http_exceptions import *


ratings_url = "http://localhost:8079/ratings/"  # URL de l'API ratings (golang)


def get_rating(id):
    response = requests.request(method="GET", url=ratings_url+id)
    return response.json(), response.status_code


def create_rating(rating_create):
    # on récupère le schéma utilisateur pour la requête vers l'API ratings
    rating_schema = RatingSchema().loads(json.dumps(rating_register), unknown=EXCLUDE)

    # on crée l'utilisateur côté API ratings
    response = requests.request(method="POST", url=ratings_url, json=rating_schema)
    if response.status_code != 201:
        return response.json(), response.status_code

    return response.json(), response.status_code


def modify_rating(id_rating, rating_update):
    """
        # on vérifie que l'utilisateur modifie ses ratings et pas ceux des autres
        if id_user != current_user.id:
            raise Forbidden
    """
    # s'il y a quelque chose à changer côté API
    rating_schema = RatingSchema().loads(json.dumps(rating_update), unknown=EXCLUDE)
    response = None
    if not RatingSchema.is_empty(rating_schema):
        # on lance la requête de modification
        response = requests.request(method="PUT", url=ratings_url+id, json=rating_schema)
        if response.status_code != 200:
            return response.json(), response.status_code

    return (response.json(), response.status_code) if response else get_rating(id)

def delete_rating(id_rating):
    """
    # on vérifie que l'utilisateur modifie ses ratings et pas ceux des autres
    if id_user != current_user.id:
        raise Forbidden
    """
    response = requests.request(method="DELETE", url=ratings_url+id_rating)
    return response.json(), response.status_code