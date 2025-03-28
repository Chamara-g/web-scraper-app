import axios from 'axios';

export const scraperAPI = async ({ siteURL }: any) => {
  try {
    const { data } = await axios({
      method: 'get',
      url: `${process.env.REACT_APP_API_HOST}/scrape?url=${siteURL}`,
      headers: {
        'Content-Type': 'application/json',
      },
    });
    return data;
  } catch (error: any) {
    if (error.response) {
      return error.response.data; // => the response payload
    } else {
      //console.log(error);
    }
  }
};
