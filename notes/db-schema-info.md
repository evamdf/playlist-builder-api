
# Chinook database info

## Tables

  table_name   |    column_name    |          data_type
---------------|-------------------| -----------------------------
 Album         | AlbumId           | integer
 Album         | Title             | character varying
 Album         | ArtistId          | integer
 Artist        | ArtistId          | integer
 Artist        | Name              | character varying
 Customer      | CustomerId        | integer
 Customer      | FirstName         | character varying
 Customer      | LastName          | character varying
 Customer      | Company           | character varying
 Customer      | Address           | character varying
 Customer      | City              | character varying
 Customer      | State             | character varying
 Customer      | Country           | character varying
 Customer      | PostalCode        | character varying
 Customer      | Phone             | character varying
 Customer      | Fax               | character varying
 Customer      | Email             | character varying
 Customer      | SupportRepId      | integer
 Employee      | EmployeeId        | integer
 Employee      | LastName          | character varying
 Employee      | FirstName         | character varying
 Employee      | Title             | character varying
 Employee      | ReportsTo         | integer
 Employee      | BirthDate         | timestamp without time zone
 Employee      | HireDate          | timestamp without time zone
 Employee      | Address           | character varying
 Employee      | City              | character varying
 Employee      | State             | character varying
 Employee      | Country           | character varying
 Employee      | PostalCode        | character varying
 Employee      | Phone             | character varying
 Employee      | Fax               | character varying
 Employee      | Email             | character varying
 Genre         | GenreId           | integer
 Genre         | Name              | character varying
 Invoice       | InvoiceId         | integer
 Invoice       | CustomerId        | integer
 Invoice       | InvoiceDate       | timestamp without time zone
 Invoice       | BillingAddress    | character varying
 Invoice       | BillingCity       | character varying
 Invoice       | BillingState      | character varying
 Invoice       | BillingCountry    | character varying
 Invoice       | BillingPostalCode | character varying
 Invoice       | Total             | numeric
 InvoiceLine   | InvoiceLineId     | integer
 InvoiceLine   | InvoiceId         | integer
 InvoiceLine   | TrackId           | integer
 InvoiceLine   | UnitPrice         | numeric
 InvoiceLine   | Quantity          | integer
 MediaType     | MediaTypeId       | integer
 MediaType     | Name              | character varying
 Playlist      | PlaylistId        | integer
 Playlist      | Name              | character varying
 PlaylistTrack | PlaylistId        | integer
 PlaylistTrack | TrackId           | integer
 Track         | TrackId           | integer
 Track         | Name              | character varying
 Track         | AlbumId           | integer
 Track         | MediaTypeId       | integer
 Track         | GenreId           | integer
 Track         | Composer          | character varying
 Track         | Milliseconds      | integer
 Track         | Bytes             | integer
 Track         | UnitPrice         | numeric

## Primary Keys

  table_name   |  column_name  
---------------|---------------
 Album         | AlbumId
 Artist        | ArtistId
 Customer      | CustomerId
 Employee      | EmployeeId
 Genre         | GenreId
 Invoice       | InvoiceId
 InvoiceLine   | InvoiceLineId
 MediaType     | MediaTypeId
 Playlist      | PlaylistId
 PlaylistTrack | PlaylistId
 PlaylistTrack | TrackId
 Track         | TrackId

## Foreign Keys

  table_name   | column_name  | foreign_table | foreign_column
---------------|--------------|---------------|----------------
 Album         | ArtistId     | Artist        | ArtistId
 Customer      | SupportRepId | Employee      | EmployeeId
 Employee      | ReportsTo    | Employee      | EmployeeId
 Invoice       | CustomerId   | Customer      | CustomerId
 InvoiceLine   | InvoiceId    | Invoice       | InvoiceId
 InvoiceLine   | TrackId      | Track         | TrackId
 PlaylistTrack | PlaylistId   | Playlist      | PlaylistId
 PlaylistTrack | TrackId      | Track         | TrackId
 Track         | AlbumId      | Album         | AlbumId
 Track         | GenreId      | Genre         | GenreId
 Track         | MediaTypeId  | MediaType     | MediaTypeId
