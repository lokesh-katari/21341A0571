import { useEffect, useState } from "react";
import "./App.css";
import axios from "axios";

function App() {
  const [product, setProduct] = useState<
    { productName: string; price: string; rating: string }[]
  >([]);

  useEffect(() => {
    const fetchProducts = async () => {
      const { data } = await axios.get(
        "http://localhost:3000/products?company=AMZ&category=Phone&top=10&minPrice=100&maxPrice=100000"
      );
      setProduct(data.products);
    };
    fetchProducts();
  }, []);

  return (
    <>
      <h1 className="text-3xl font-bold underline">
        Hello world! these are Products
      </h1>

      <div className="flex  items-center just gap-5">
        {product.map((item, index) => (
          <a
            href="#"
            className="block max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
          >
            <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
              {item.productName}
            </h5>
            <p className="font-normal text-gray-700 dark:text-gray-400">
              {item.price}
            </p>
            <p className="font-normal text-gray-700 dark:text-gray-400">
              {item.rating}
            </p>
          </a>
        ))}
      </div>
    </>
  );
}

export default App;
