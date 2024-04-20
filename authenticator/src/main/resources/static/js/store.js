/*jshint eqeqeq:false */
(function (window) {
	'use strict';

	/**
	 * Creates a new client side storage object and will create an empty
	 * collection if no collection already exists.
	 *
	 * @param {string} name The name of our DB we want to use
	 * @param {function} callback Our fake DB uses callbacks because in
	 * real life you probably would be making AJAX calls
	 */
	function Store(name, callback) {
		callback = callback || function () {};

		this._dbName = name;

		if (!localStorage.getItem(name)) {
			var todos = [];

			localStorage.setItem(name, JSON.stringify(todos));
		}

		callback.call(this, JSON.parse(localStorage.getItem(name)));
	}

	/**
	 * Finds items based on a query given as a JS object
	 *
	 * @param {object} query The query to match against (i.e. {foo: 'bar'})
	 * @param {function} callback	 The callback to fire when the query has
	 * completed running
	 *
	 * @example
	 * db.find({foo: 'bar', hello: 'world'}, function (data) {
	 *	 // data will return any items that have foo: bar and
	 *	 // hello: world in their properties
	 * });
	 */
	Store.prototype.find = function (query, callback) {
		if (!callback) {
			return;
		}

		const xhr = new XMLHttpRequest();
		var s = '';
		for (var q in query){
			if (s === ''){
				s = '?' + q + '=' + query[q]
			} else {
				s = s + '&' + q + '=' + query[q]
			}
		}
		xhr.open('GET', 'http://localhost:8080/task' + s);
		xhr.responseType = "json";
		xhr.onreadystatechange = function() {
			if (xhr.readyState !== 4 || xhr.status !== 200) {
				return;
			}
			const response = xhr.response;
			callback.call(this, response);
		}
		xhr.send();

		// var todos = JSON.parse(localStorage.getItem(this._dbName));
		//
		// callback.call(this, todos.filter(function (todo) {
		// 	for (var q in query) {
		// 		if (query[q] !== todo[q]) {
		// 			return false;
		// 		}
		// 	}
		// 	return true;
		// }));
	};

	/**
	 * Will retrieve all data from the collection
	 *
	 * @param {function} callback The callback to fire upon retrieving data
	 */
	Store.prototype.findAll = function (callback) {
		callback = callback || function () {};
		const xhr = new XMLHttpRequest();
		xhr.open('GET', 'http://localhost:8080/task');
		xhr.responseType = "json";
		xhr.onreadystatechange = function() {
			if (xhr.readyState !== 4 || xhr.status !== 200) {
				return;
			}
			const response = xhr.response;
			callback.call(this, response);
		}
		xhr.send();
		// callback.call(this, JSON.parse(localStorage.getItem(this._dbName)));
	};

	/**
	 * Will save the given data to the DB. If no item exists it will create a new
	 * item, otherwise it'll simply update an existing item's properties
	 *
	 * @param {object} updateData The data to save back into the DB
	 * @param {function} callback The callback to fire after saving
	 * @param {number} id An optional param to enter an ID of an item to update
	 */
	Store.prototype.save = function (updateData, callback, id) {
		var todos = JSON.parse(localStorage.getItem(this._dbName));

		callback = callback || function() {};

		// If an ID was actually given, find the item and update each property
		if (id) {
			updateData.id = id;
			const xhr = new XMLHttpRequest();
			xhr.open('PUT', 'http://localhost:8080/task');
			xhr.setRequestHeader("Content-Type", "application/json");
			xhr.onreadystatechange = function() {
				if (xhr.readyState !== 4 || xhr.status !== 200) {
					return;
				}
				// Store.prototype.findAll(callback);
				callback.call(this, [updateData]);
			}
			xhr.send(JSON.stringify(updateData));
			// for (var i = 0; i < todos.length; i++) {
			// 	if (todos[i].id === id) {
			// 		for (var key in updateData) {
			// 			todos[i][key] = updateData[key];
			// 		}
			// 		break;
			// 	}
			// }
			//
			// localStorage.setItem(this._dbName, JSON.stringify(todos));
			// callback.call(this, updateData);
		} else {
			// Generate an ID
			const xhr = new XMLHttpRequest();
			xhr.open('POST', 'http://localhost:8080/task');
			xhr.responseType = "json"
			xhr.setRequestHeader("Content-Type", "application/json");
			xhr.onreadystatechange = function() {
				if (xhr.readyState !== 4 || xhr.status !== 200) {
					return;
				}
				const response = xhr.response;
				// Store.prototype.findAll(callback);
				callback.call(this, [response]);
			}
			xhr.send(JSON.stringify(updateData));
			// updateData.id = new Date().getTime();
			//
			// todos.push(updateData);
			// localStorage.setItem(this._dbName, JSON.stringify(todos));
			// callback.call(this, [updateData]);
		}
	};

	/**
	 * Will remove an item from the Store based on its ID
	 *
	 * @param {number} id The ID of the item you want to remove
	 * @param {function} callback The callback to fire after saving
	 */
	Store.prototype.remove = function (id, callback) {
		const xhr = new XMLHttpRequest();
		xhr.open('DELETE', 'http://localhost:8080/task?id=' + id);
		xhr.onreadystatechange = function() {
			if (xhr.readyState !== 4 || xhr.status !== 200) {
				return;
			}
			Store.prototype.findAll(callback);
		}
		xhr.send();
		// var todos = JSON.parse(localStorage.getItem(this._dbName));
		//
		// for (var i = 0; i < todos.length; i++) {
		// 	if (todos[i].id == id) {
		// 		todos.splice(i, 1);
		// 		break;
		// 	}
		// }
		//
		// localStorage.setItem(this._dbName, JSON.stringify(todos));
		// callback.call(this, todos);
	};

	/**
	 * Will drop all storage and start fresh
	 *
	 * @param {function} callback The callback to fire after dropping the data
	 */
	Store.prototype.drop = function (callback) {
		const xhr = new XMLHttpRequest();
		xhr.open('DELETE', 'http://localhost:8080/task');
		xhr.onreadystatechange = function() {
			if (xhr.readyState !== 4 || xhr.status !== 200) {
				return;
			}
			callback.call(this, []);
		}
		xhr.send();
		// var todos = [];
		// localStorage.setItem(this._dbName, JSON.stringify(todos));
		// callback.call(this, todos);
	};

	// Export to window
	window.app = window.app || {};
	window.app.Store = Store;
})(window);
