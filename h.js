db.geekFlareCollection.updateOne({"product" : "bottles"}, {$set : {"Qty": 40}})

db.users.update({"username": "pedro"}, {$set: {"age": 55}})

console.log("hola");

// delete
db.users.deleteMany({ username: "mario" })
