db = db.getSiblingDB("cluster");

db.users.insertMany([
  { name: "Alice", email: "alice@example.com" },
  { name: "testing", email: "bob@example.com" }
]);

db.products.insertOne({ name: "Widget", price: 9.99 });