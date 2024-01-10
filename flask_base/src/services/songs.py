import json
import requests
from flask import jsonify
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from src.schemas.song import SongSchema
from src.schemas.user import UserSchema
from src.models.user import User as UserModel
from src.models.http_exceptions import *
import src.repositories.users as users_repository


songs_url = "http://localhost:8081/songs/"  # URL de l'API users (golang)

def get_song(id):
    print("debug0")
    response = requests.request(method="GET", url=songs_url+id)
    if response.status_code != 200 :
        return jsonify({'error': 'Failed to update rating'}), response.status_code
    else:
        return jsonify({'message': 'Rating updated successfully'}), 200

def get_songs():
    print("debug0")
    response = requests.request(method="GET", url=songs_url)
    print(response.status_code)
    return response.json(), response.status_code

def delete_song(id):
    """
    # on vérifie que l'utilisateur modifie ses ratings et pas ceux des autres
    if id_user != current_user.id:
        raise Forbidden
    """
    response = requests.request(method="DELETE", url=songs_url+id)
    if response.status_code != 204:
        return jsonify({'error': 'Failed to delete rating'}), response.status_code
    else:
        return jsonify({'message': 'Rating deleted successfully'}), 204 

def create_song(song):

    # on récupère le schéma utilisateur pour la requête vers l'API users
    # on crée l'utilisateur côté API users

    song_schema = SongSchema().loads(json.dumps(song, default=str), unknown=EXCLUDE)
    print(song_schema)

    response = requests.post(songs_url, json=song)


    return response.json(), response.status_code

def update_song(id,song):
    song_schema = SongSchema().loads(json.dumps(song, default=str), unknown=EXCLUDE)
    print(song_schema)
    response = requests.put(songs_url+id, json=song)
    print(response.status_code)
    if response.status_code != 200:
        return jsonify({'error': 'Failed to update rating'}), response.status_code
    else:
        return jsonify({'message': 'Rating updated successfully'}), 200