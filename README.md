# GPS-Tracking
Clone or unzip project directory.

    git clone https://github.com/pudg/GPS-Tracking.git

## Backend
### Install
Step into the backend directory of the project.

    cd GPS-tracking/backend

Install required packages.

    go mod tidy

Launch the backend server.

    go run .

## Frontend
### Install
Open a separate terminal window and step into frontend directory of the project.

    cd GPS-tracking/frontend

Install required packages using `npm`

    npm install

Launch the frontend 

    npm run dev

## Use
Visit [http://localhost:3000/](http://localhost:3000/)

Register an account by clicking on the `Register` button on the home page.

Once registered, you will be redirected to the tracking page. 
 * `/tracking` requires authentication so visiting the page without loging in will redirect the user to `/login`.

The `/tracking` page has a basic layout with the map on the right-hand side, and device list on the left-hand side.

The settings menu on the top left-hand side allows the user to:
 * `Query` the _OneStep_ API to get a list of all devices.
 * `Track` all devices on the map.
 * `Sort` the devices by _device model_.
 * `Save` all setting preferences and uploaded files to database.

To inspect the database after any `save` operation visit [http://localhost:8000/api/database](http://localhost:8000/api/database)

 * Only included for coding challenge transparency, this wouldn't make it in release.

## API
View Postman generated [API documentation.](https://documenter.getpostman.com/view/29003440/2s9YRGwoae)