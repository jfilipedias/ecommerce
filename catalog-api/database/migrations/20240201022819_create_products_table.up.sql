CREATE TABLE products (
  id VARCHAR(36) NOT NULL,
  name VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  category_id VARCHAR(36) NOT NULL,
  image_url VARCHAR(255) NOT NULL,
  PRIMARY KEY(id),
  CONSTRAINT fk_category
   FOREIGN KEY(category_id) 
      REFERENCES categories(id)
)