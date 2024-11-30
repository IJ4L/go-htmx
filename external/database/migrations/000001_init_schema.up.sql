CREATE TABLE product (
  product_id SERIAL PRIMARY KEY,              
  name VARCHAR(255) NOT NULL,                 
  category VARCHAR(100),                      
  price DECIMAL(10, 2) NOT NULL,              
  image_url VARCHAR(255),                     
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  
);


CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;  
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON product
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
