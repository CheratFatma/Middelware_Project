from flask import Flask, request, jsonify
import requests

app = Flask(__name__)

# URL de votre API Go
API_GO_URL = "http://localhost:8080"  

#API USERS
# Endpoint pour récupérer la liste des utilisateurs
@app.route('/api/users', methods=['GET'])
def get_users():
    try:
        # Appel à l'API Go pour récupérer la liste des utilisateurs
        response = requests.get(f"{API_GO_URL}/users")

        # Vérifiez si la requête a réussi (code HTTP 200)
        if response.status_code == 200:
            # Renvoyez les données de l'API Go comme réponse JSON dans votre API Flask
            return jsonify(response.json())
        else:
            # Gérez les erreurs si la requête a échoué
            return jsonify({'error': 'Failed to retrieve users'}), 500

    except Exception as e:
        return jsonify({'error': str(e)}), 500

# Endpoint pour récupérer un utilisateur par son ID
@app.route('/api/users/<int:user_id>', methods=['GET'])
def get_user_by_id(user_id):
    try:
        # Appel à l'API Go pour récupérer un utilisateur par son ID
        response = requests.get(f"{API_GO_URL}/users/{user_id}")

        # Vérifiez si la requête a réussi (code HTTP 200)
        if response.status_code == 200:
            # Renvoyez les données de l'API Go comme réponse JSON dans votre API Flask
            return jsonify(response.json())
        elif response.status_code == 404:
            # Gérez le cas où l'utilisateur n'est pas trouvé
            return jsonify({'error': 'User not found'}), 404
        else:
            # Gérez les autres erreurs
            return jsonify({'error': 'Failed to retrieve user'}), 500

    except Exception as e:
        return jsonify({'error': str(e)}), 500

# Endpoint pour ajouter un nouvel utilisateur
@app.route('/api/users', methods=['POST'])
def add_new_user():
    try:
        # Récupérez les données du corps de la requête
        user_data = request.get_json()

        # Appel à l'API Go pour ajouter un nouvel utilisateur
        response = requests.post(f"{API_GO_URL}/users", json=user_data)

        # Vérifiez si la requête a réussi (code HTTP 201)
        if response.status_code == 201:
            return jsonify({'message': 'New user added successfully'}), 201
        else:
            # Gérez les erreurs si la requête a échoué
            return jsonify({'error': 'Failed to add new user'}), 500

    except Exception as e:
        return jsonify({'error': str(e)}), 500

# Endpoint pour mettre à jour un utilisateur par son ID
@app.route('/api/users/<int:user_id>', methods=['PUT'])
def update_user_by_id(user_id):
    try:
        # Récupérez les données du corps de la requête
        updated_user_data = request.get_json()

        # Appel à l'API Go pour mettre à jour un utilisateur par son ID
        response = requests.put(f"{API_GO_URL}/users/{user_id}", json=updated_user_data)

        # Vérifiez si la requête a réussi (code HTTP 200)
        if response.status_code == 200:
            return jsonify({'message': 'User updated successfully'}), 200
        elif response.status_code == 404:
            # Gérez le cas où l'utilisateur n'est pas trouvé
            return jsonify({'error': 'User not found'}), 404
        else:
            # Gérez les autres erreurs
            return jsonify({'error': 'Failed to update user'}), 500

    except Exception as e:
        return jsonify({'error': str(e)}), 500

# Endpoint pour supprimer un utilisateur par son ID
@app.route('/api/users/<int:user_id>', methods=['DELETE'])
def delete_user_by_id(user_id):
    try:
        # Appel à l'API Go pour supprimer un utilisateur par son ID
        response = requests.delete(f"{API_GO_URL}/users/{user_id}")

        # Vérifiez si la requête a réussi (code HTTP 200)
        if response.status_code == 200:
            return jsonify({'message': 'User deleted successfully'}), 200
        elif response.status_code == 404:
            # Gérez le cas où l'utilisateur n'est pas trouvé
            return jsonify({'error': 'User not found'}), 404
        else:
            # Gérez les autres erreurs
            return jsonify({'error': 'Failed to delete user'}), 500

    except Exception as e:
        return jsonify({'error': str(e)}), 500

#API SONGS
# Endpoint pour récupérer la liste des chansons
@app.route('/api/songs', methods=['GET'])
def get_songs():
    try:
        # Appel à l'API Go pour récupérer la liste des chansons
        response = requests.get(f"{API_GO_URL}/songs")

        # Vérifiez si la requête a réussi (code HTTP 200)
        if response.status_code == 200:
            # Renvoyez les données de l'API Go comme réponse JSON dans votre API Flask
            return jsonify(response.json())
        else:
            # Gérez les erreurs si la requête a échoué
            return jsonify({'error': 'Failed to retrieve songs'}), 500

    except Exception as e:
        return jsonify({'error': str(e)}), 500

# Endpoint pour récupérer une chanson par son ID
@app.route('/api/songs/<int:song_id>', methods=['GET'])
def get_song_by_id(song_id):
    try:
        # Appel à l'API Go pour récupérer une chanson par son ID
        response = requests.get(f"{API_GO_URL}/songs/{song_id}")

        # Vérifiez si la requête a réussi (code HTTP 200)
        if response.status_code == 200:
            # Renvoyez les données de l'API Go comme réponse JSON dans votre API Flask
            return jsonify(response.json())
        elif response.status_code == 404:
            # Gérez le cas où la chanson n'est pas trouvée
            return jsonify({'error': 'Song not found'}), 404
        else:
            # Gérez les autres erreurs
            return jsonify({'error': 'Failed to retrieve song'}), 500

    except Exception as e:
        return jsonify({'error': str(e)}), 500

# Endpoint pour ajouter une nouvelle chanson
@app.route('/api/songs', methods=['POST'])
def add_new_song():
    try:
        # Récupérez les données du corps de la requête
        song_data = request.get_json()

        # Appel à l'API Go pour ajouter une nouvelle chanson
        response = requests.post(f"{API_GO_URL}/songs", json=song_data)

        # Vérifiez si la requête a réussi (code HTTP 201)
        if response.status_code == 201:
            return jsonify({'message': 'New song added successfully'}), 201
        else:
            # Gérez les erreurs si la requête a échoué
            return jsonify({'error': 'Failed to add new song'}), 500

    except Exception as e:
        return jsonify({'error': str(e)}), 500

# Endpoint pour mettre à jour une chanson par son ID
@app.route('/api/songs/<int:song_id>', methods=['PUT'])
def update_song_by_id(song_id):
    try:
        # Récupérez les données du corps de la requête
        updated_song_data = request.get_json()

        # Appel à l'API Go pour mettre à jour une chanson par son ID
        response = requests.put(f"{API_GO_URL}/songs/{song_id}", json=updated_song_data)

        # Vérifiez si la requête a réussi (code HTTP 200)
        if response.status_code == 200:
            return jsonify({'message': 'Song updated successfully'}), 200
        elif response.status_code == 404:
            # Gérez le cas où la chanson n'est pas trouvée
            return jsonify({'error': 'Song not found'}), 404
        else:
            # Gérez les autres erreurs
            return jsonify({'error': 'Failed to update song'}), 500

    except Exception as e:
        return jsonify({'error': str(e)}), 500

# Endpoint pour supprimer une chanson par son ID
@app.route('/api/songs/<int:song_id>', methods=['DELETE'])
def delete_song_by_id(song_id):
    try:
        # Appel à l'API Go pour supprimer une chanson par son ID
        response = requests.delete(f"{API_GO_URL}/songs/{song_id}")

        # Vérifiez si la requête a réussi (code HTTP 200)
        if response.status_code == 200:
            return jsonify({'message': 'Song deleted successfully'}), 200
        elif response.status_code == 404:
            # Gérez le cas où la chanson n'est pas trouvée
            return jsonify({'error': 'Song not found'}), 404
        else:
            # Gérez les autres erreurs
            return jsonify({'error': 'Failed to delete song'}), 500

    except Exception as e:
        return jsonify({'error': str(e)}), 500

# Exécutez l'application Flask si le script est exécuté directement
if __name__ == '__main__':
    app.run(debug=True)


