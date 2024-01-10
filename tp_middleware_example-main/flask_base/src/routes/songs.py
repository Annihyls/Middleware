import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from src.models.http_exceptions import *
from src.schemas.song import SongUpdateSchema
from src.schemas.song import SongCreateSchema
from src.schemas.errors import *
import src.services.songs as songs_service

# from routes import songs
songs = Blueprint(name="songs", import_name=__name__)


@songs.route('/<id>', methods=['GET'])
#@login_required
def get_song(id):
    print("debug2")
    """
    ---
    get:
      description: Getting a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '201':
          description: Ok
          content:
            application/json:
              schema: song
            application/yaml:
              schema: song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
    """
    return songs_service.get_song(id)

@songs.route('/', methods=['GET'])
#@login_required
def get_songs():
    print("debug2")
    """
    ---
    get:
      description: Getting a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '201':
          description: Ok
          content:
            application/json:
              schema: song
            application/yaml:
              schema: song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
    """
    return songs_service.get_songs()

@songs.route('/<id>', methods=['DELETE'])
#@login_required
def delete_song(id):
    print("debug2")
    """
    ---
    get:
      description: Getting a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '201':
          description: Ok
          content:
            application/json:
              schema: song
            application/yaml:
              schema: song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
    """
    return songs_service.delete_song(id)


@songs.route('/<id>', methods=['PUT'])
#@login_required
def put_song(id):
    """
    ---
    put:
      description: Updating a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      requestBody:
        required: true
        content:
            application/json:
                schema: songUpdate
      responses:
        '201':
          description: Ok
          content:
            application/json:
              schema: song
            application/yaml:
              schema: song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - songs
    """
    # parser le body
    try:
        song_update = SongUpdateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # modification de l'utilisateur (songname, nom, mot de passe, etc.)
    try:
        return songs_service.update_song(id, song_update)
    except Conflict:
        error = ConflictSchema().loads(json.dumps({"message": "song already exists"}))
        return error, error.get("code")
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return error, error.get("code")
    except Forbidden:
        error = ForbiddenSchema().loads(json.dumps({"message": "Can't manage other songs"}))
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")

@songs.route('/', methods=['POST'])
def create_song():
    """
    ---
    create:
      description: create a song
      parameters:
      requestBody:
        required: true
        content:
            application/json:
                schema: SongUpdate
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - songs
    """
    # parser le body
    try:
        song = SongCreateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # modification de l'utilisateur (title, artist, etc.)
    try:
        return songs_service.create_song(song)
    except Conflict:
        error = ConflictSchema().loads(json.dumps({"message": "Song already exists"}))
        return error, error.get("code")
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return error, error.get("code")
    except Forbidden:
        error = ForbiddenSchema().loads(json.dumps({"message": "Can't manage other songs"}))
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")