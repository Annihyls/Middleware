import json
from flask import jsonify
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

def delete_rating(id_rating):
    """
    # on vérifie que l'utilisateur modifie ses ratings et pas ceux des autres
    if id_user != current_user.id:
        raise Forbidden
    """
    response = requests.request(method="DELETE", url=ratings_url+id_rating)
    if response.status_code != 204:
        return jsonify({'error': 'Failed to delete rating'}), response.status_code
    else:
        return jsonify({'message': 'Rating deleted successfully'}), 204



def create_rating(rating_create):
    # on récupère le schéma rating pour la requête vers l'API ratings
    print("Le problème se situe ci-dessous !!!")
    rating_schema = RatingSchema().loads(json.dumps(rating_create), unknown=EXCLUDE)
    print("-------C'EST BON C'EST PASSE-------")
    # on crée l'utilisateur côté API ratings
    response = requests.request(method="POST", url=ratings_url, json=rating_schema)
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
        response = requests.request(method="PUT", url=ratings_url+id_rating, json=rating_schema)
        if response.status_code != 200:
            return jsonify({'error': 'Failed to update rating'}), response.status_code
        else:
            return jsonify({'message': 'Rating updated successfully'}), 200
    else:
        # Je met ça car il y a un bug des fois flask n'accepte pas la requete. Il faut changer son update
        return jsonify({'bug': 'Change your update, because it\'s buggy'}), 400

