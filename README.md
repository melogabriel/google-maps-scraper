# Google maps scraper
![build](https://github.com/gosom/google-maps-scraper/actions/workflows/build.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/gosom/google-maps-scraper)](https://goreportcard.com/report/github.com/gosom/google-maps-scraper)
[![Image of repositories-views-counter](https://github.com/melogabriel/repositories-views-counter/blob/master/svg/946732374/badge.svg)](https://github.com/melogabriel/repositories-views-counter/blob/master/readme/946732374/week.md)

> A free and open-source Google Maps scraper with both command line and web UI options. This tool is easy to use and allows you to extract data from Google Maps efficiently.

[Google Maps Scraper](https://google-maps-scraper-6lcb.onrender.com) is a web-based tool designed to extract detailed business information from Google Maps based on user-defined search parameters. This deployment makes it easier to use compared to the previous fork version, as it is now accessible directly via a web interface without requiring local setup.

## Features

- **Keyword-Based Search**: Input specific keywords to target businesses or places of interest.
- **Language Selection**: Choose the language for the search results to match regional preferences.
- **Location Settings**:
  - **Zoom**: Adjust the zoom level to refine the search area.
  - **Latitude and Longitude**: Specify geographical coordinates to center the search.
- **Advanced Options**:
  - **Fast Mode (BETA)**: Enable faster scraping with potential trade-offs in data accuracy.
  - **Radius (BETA)**: Define a specific radius for the search area.
  - **Depth**: Set the depth of search to control the number of results.
  - **Fetch Emails**: Attempt to retrieve email addresses associated with the businesses.
  - **Max Job Time**: Limit the maximum time allocated for a scraping task.
- **Proxy Support**: Configure proxies to manage request routing and avoid potential restrictions.

## Getting Started

How scrape Google Maps data:

1. **Access the Application**: Navigate to the [Google Maps Scraper](https://google-maps-scraper-6lcb.onrender.com) web application.
2. **Set Up a New Scraping Job**:
   - **Job Name**: Assign a unique name to your scraping task.
   - **Keywords**: Enter the search terms relevant to your target data.
   - **Language**: Select the desired language for the search results.
   - **Location Settings**: Adjust the zoom level and specify latitude and longitude coordinates to define the search area.
   - **Advanced Options**: Configure additional settings such as fast mode, search radius, depth, email fetching, and maximum job time.
   - **Proxies**: If necessary, input proxy details to manage request routing.
3. **Initiate Scraping**: Click the "Start Scraping" button to commence the data extraction process.
4. **Monitor Job Status**: View the status of your scraping jobs in the "Job Details" section, which provides information on job ID, name, date, status, and available actions.

## Why This Deployment?

This deployment improves accessibility and ease of use compared to the previous forked version, eliminating the need for local installation or configuration. Users can now run scraping jobs directly from the web without additional setup.

## What Google maps scraper does

Web based google maps scraper build using 

[scrapemate](https://github.com/gosom/scrapemate) web crawling framework.

You can use this repository either as is, or you can use its code as a base and
customize it to your needs

![Example GIF](img/example.gif)

### REST API
The Google Maps Scraper provides a RESTful API for programmatic management of scraping tasks.

### Key Endpoints

- POST /api/v1/jobs: Create a new scraping job
- GET /api/v1/jobs: List all jobs
- GET /api/v1/jobs/{id}: Get details of a specific job
- DELETE /api/v1/jobs/{id}: Delete a job
- GET /api/v1/jobs/{id}/download: Download job results as CSV

For detailed API documentation, refer to the OpenAPI 3.0.3 specification available through Swagger UI or Redoc when running the app https://localhost:8080/api/docs


## 🌟 Support the Project!

If you find this tool useful, consider giving it a **star** on GitHub. 
Your support helps ensure continued improvement and maintenance.

## Parameters

- Extracts many data points from google maps
- Exports the data to CSV, JSON or PostgreSQL 
- Performance about 120 urls per minute (-depth 1 -c 8)
- Extendable to write your own exporter
- Dockerized for easy run in multiple platforms
- Scalable in multiple machines
- Optionally extracts emails from the website of the business
- SOCKS5/HTTP/HTTPS proxy support
- Serverless execution via AWS Lambda functions (experimental & no documentation yet)
- Fast Mode (BETA)

## Notes on email extraction

By default email extraction is disabled. 

If you enable email extraction (see quickstart) then the scraper will visit the 
website of the business (if exists) and it will try to extract the emails from the
page.

For the moment it only checks only one page of the website (the one that is registered in Gmaps). At some point, it will be added support to try to extract from other pages like about, contact, impressum etc. 


Keep in mind that enabling email extraction results to larger processing time, since more
pages are scraped. 

## Fast Mode

Fast mode returns you at most 21 search results per query ordered by distance from the **latitude** and **longitude** provided.
All the results are within the specified **radius**

It does not contain all the data points but basic ones. 
However it provides the ability to extract data really fast. 

When you use the fast mode ensure that you have provided:
- zoom
- radius (in meters)
- latitude
- longitude


**Fast mode is Beta, you may experience blocking**

## Extracted Data Points

#### 1. `input_id`
- Internal identifier for the input query.

#### 2. `link`
- Direct URL to the business listing on Google Maps.

#### 3. `title`
- Name of the business.

#### 4. `category`
- Business type or category (e.g., Restaurant, Hotel).

#### 5. `address`
- Street address of the business.

#### 6. `open_hours`
- Business operating hours.

#### 7. `popular_times`
- Estimated visitor traffic at different times of the day.

#### 8. `website`
- Official business website.

#### 9. `phone`
- Business contact phone number.

#### 10. `plus_code`
- Shortcode representing the precise location of the business.

#### 11. `review_count`
- Total number of customer reviews.

#### 12. `review_rating`
- Average star rating based on reviews.

#### 13. `reviews_per_rating`
- Breakdown of reviews by each star rating (e.g., number of 5-star, 4-star reviews).

#### 14. `latitude`
- Latitude coordinate of the business location.

#### 15. `longitude`
- Longitude coordinate of the business location.

#### 16. `cid`
- **Customer ID** (CID) used by Google Maps to uniquely identify a business listing. This ID remains stable across updates and can be used in URLs.
- **Example:** `3D3174616216150310598`

#### 17. `status`
- Business status (e.g., open, closed, temporarily closed).

#### 18. `descriptions`
- Brief description of the business.

#### 19. `reviews_link`
- Direct link to the reviews section of the business listing.

#### 20. `thumbnail`
- URL to a thumbnail image of the business.

#### 21. `timezone`
- Time zone of the business location.

#### 22. `price_range`
- Price range of the business (`$`, `$$`, `$$$`).

#### 23. `data_id`
- An internal Google Maps identifier composed of two hexadecimal values separated by a colon.
- **Structure:** `<spatial_hex>:<listing_hex>`
- **Example:** `0x3eb33fecd7dfa167:0x2c0e80a0f5d57ec6`
- **Note:** This value may change if the listing is updated and should not be used for permanent identification.

#### 24. `images`
- Links to images associated with the business.

#### 25. `reservations`
- Link to book reservations (if available).

#### 26. `order_online`
- Link to place online orders.

#### 27. `menu`
- Link to the menu (for applicable businesses).

#### 28. `owner`
- Indicates whether the business listing is claimed by the owner.

#### 29. `complete_address`
- Fully formatted address of the business.

#### 30. `about`
- Additional information about the business.

#### 31. `user_reviews`
- Collection of customer reviews, including text, rating, and timestamp.

#### 32. `emails`
- Email addresses associated with the business, if available.

**Note**: email is empty by default (see Usage)

**Note**: Input id is an ID that you can define per query. By default it's a UUID
In order to define it you can have an input file like:

```
Matsuhisa Athens #!#MyIDentifier
```

## References

For more instruction you may also read the following links

- https://blog.gkomninos.com/how-to-extract-data-from-google-maps-using-golang
- https://blog.gkomninos.com/distributed-google-maps-scraping
- https://github.com/omkarcloud/google-maps-scraper/tree/master (also a nice project) [many thanks for the idea to extract the data by utilizing the JS objects]


## Licence

This code is licensed under the MIT License


## Contributing

Please open an ISSUE or make a Pull Request


Thank you for considering support for the project. Every bit of assistance helps maintain momentum and enhances the scraper’s capabilities!



