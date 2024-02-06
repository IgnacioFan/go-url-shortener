import axios from "axios";

axios.defaults.baseURL = "http://localhost";

const createShortUrl = async (longUrl: string) => {
  try {
    const response = await axios.post("/v1/urls", {
      headers: { "content-type": "application/json" },
      long_url: longUrl,
    });
    return response.data
  } catch (err) {
    return Promise.reject(err)
  }
}

export { createShortUrl }
