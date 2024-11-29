# Groupie Tracker

## Objectives

Groupie Tracker is a project focused on creating a user-friendly website by interacting with a given API and displaying its data in an engaging and visual way. The project emphasizes data manipulation, visualization, and client-server interaction.

The API consists of four main parts:

1. **Artists**: Contains information about bands and artists, including:

   - Names
   - Images
   - Year they began their activity
   - Date of their first album
   - Members

2. **Locations**: Details their last and/or upcoming concert locations.

3. **Dates**: Provides information on their last and/or upcoming concert dates.

4. **Relation**: Links the artists, dates, and locations together.

Using this API, you will build a website to display band information through various data visualizations, such as:

- Blocks
- Cards
- Tables
- Lists
- Pages
- Graphics

The approach to displaying the information is up to you, but the goal is to create an intuitive and visually appealing user experience.

## Features

### Event/Action System

A key feature of the project is implementing **client-server interaction**, which involves:

- Creating an event or action triggered by the client, time, or another factor.
- Communicating with the server to retrieve data through a [request-response](https://en.wikipedia.org/wiki/Request%E2%80%93response) cycle.
- Visualizing the triggered event/action on the site.

## Instructions

- The backend **must** be written in **Go**.
- The site and server **cannot crash** at any time.
- All pages must work correctly, and errors should be handled gracefully.
- The code must follow **best practices**.
- It is recommended to include **unit tests** for your code.

### Allowed Packages

- Only standard Go packages are permitted.

## Usage

An example of a RESTful API can be found [here](https://example.com).

## Learning Goals

This project will help you develop skills in:

- Manipulating and storing data
- Working with JSON files and formats
- Building and styling HTML pages
- Creating and displaying events
- Understanding client-server communication
