package utils

var CreateTableQueriesSlice = []string{

	`
    CREATE TABLE IF NOT EXISTS MediaType (
        MediaTypeId SERIAL PRIMARY KEY,
        Name VARCHAR(120)
    );`,

	`
    CREATE TABLE IF NOT EXISTS Playlist (
        PlaylistId SERIAL PRIMARY KEY,
        Name VARCHAR(120)
    );`,

	`
    CREATE TABLE IF NOT EXISTS Artist (
        ArtistId SERIAL PRIMARY KEY,
        Name VARCHAR(120)
    );`,

	`
    CREATE TABLE IF NOT EXISTS Genre (
        GenreId SERIAL PRIMARY KEY,
        Name VARCHAR(120)
    );`,

	`
    CREATE TABLE IF NOT EXISTS Album (
        AlbumId SERIAL PRIMARY KEY,
        Title VARCHAR(160) NOT NULL,
        ArtistId INTEGER NOT NULL,
        FOREIGN KEY (ArtistId) REFERENCES Artist (ArtistId)
    );`,

	`
        CREATE TABLE IF NOT EXISTS Employee (
            EmployeeId SERIAL PRIMARY KEY,
            LastName VARCHAR(20) NOT NULL,
            FirstName VARCHAR(20) NOT NULL,
            Title VARCHAR(30),
            ReportsTo INTEGER,
            BirthDate TIMESTAMP,
            HireDate TIMESTAMP,
            Address VARCHAR(70),
            City VARCHAR(40),
            State VARCHAR(40),
            Country VARCHAR(40),
            PostalCode VARCHAR(10),
            Phone VARCHAR(24),
            Fax VARCHAR(24),
            Email VARCHAR(60),
            FOREIGN KEY (ReportsTo) REFERENCES Employee (EmployeeId) 
        );`,
	`
        CREATE TABLE IF NOT EXISTS Customer (
            CustomerId SERIAL PRIMARY KEY,
            FirstName VARCHAR(40) NOT NULL,
            LastName VARCHAR(20) NOT NULL,
            Company VARCHAR(80),
            Address VARCHAR(70),
            City VARCHAR(40),
            State VARCHAR(40),
            Country VARCHAR(40),
            PostalCode VARCHAR(10),
            Phone VARCHAR(24),
            Fax VARCHAR(24),
            Email VARCHAR(60) NOT NULL,
            SupportRepId INTEGER,
            FOREIGN KEY (SupportRepId) REFERENCES Employee (EmployeeId) 
        );`,
	`
        CREATE TABLE IF NOT EXISTS Invoice (
            InvoiceId SERIAL PRIMARY KEY,
            CustomerId INTEGER NOT NULL,
            InvoiceDate TIMESTAMP NOT NULL,
            BillingAddress VARCHAR(70),
            BillingCity VARCHAR(40),
            BillingState VARCHAR(40),
            BillingCountry VARCHAR(40),
            BillingPostalCode VARCHAR(10),
            Total NUMERIC(10, 2) NOT NULL,
            FOREIGN KEY (CustomerId) REFERENCES Customer (CustomerId) 
        );`,
	`
        CREATE TABLE IF NOT EXISTS Track (
            TrackId SERIAL PRIMARY KEY,
            Name VARCHAR(200) NOT NULL,
            AlbumId INTEGER,
            MediaTypeId INTEGER NOT NULL,
            GenreId INTEGER,
            Composer VARCHAR(220),
            Milliseconds INTEGER NOT NULL,
            Bytes INTEGER,
            UnitPrice NUMERIC(10, 2) NOT NULL,
            FOREIGN KEY (MediaTypeId) REFERENCES MediaType (MediaTypeId) ,
            FOREIGN KEY (GenreId) REFERENCES Genre (GenreId) ,
            FOREIGN KEY (AlbumId) REFERENCES Album (AlbumId) 
        );`,

	`
        CREATE TABLE IF NOT EXISTS PlaylistTrack (
            PlaylistId INTEGER NOT NULL,
            TrackId INTEGER NOT NULL,
            PRIMARY KEY (PlaylistId, TrackId),
            FOREIGN KEY (PlaylistId) REFERENCES Playlist (PlaylistId) ,
            FOREIGN KEY (TrackId) REFERENCES Track (TrackId) 
        );`,
	`
        CREATE TABLE IF NOT EXISTS InvoiceLine (
            InvoiceLineId SERIAL PRIMARY KEY,
            InvoiceId INTEGER NOT NULL,
            TrackId INTEGER NOT NULL,
            UnitPrice NUMERIC(10, 2) NOT NULL,
            Quantity INTEGER NOT NULL,
            FOREIGN KEY (InvoiceId) REFERENCES Invoice (InvoiceId) ,
            FOREIGN KEY (TrackId) REFERENCES Track (TrackId)
        );`,
}

// =====================================================
var CreateTableQueriesMap = map[string]string{
	"mediaType": `
        CREATE TABLE IF NOT EXISTS Playlist (
            PlaylistId SERIAL PRIMARY KEY,
            Name VARCHAR(120)
        );`,
	"playlist": `
       CREATE TABLE IF NOT EXISTS Playlist (
           PlaylistId SERIAL PRIMARY KEY,
           Name VARCHAR(120)
       );`,
	"artist": `
        CREATE TABLE IF NOT EXISTS Artist (
            ArtistId SERIAL PRIMARY KEY,
            Name VARCHAR(120)
        );`,
	"genre": `
        CREATE TABLE IF NOT EXISTS Genre (
            GenreId SERIAL PRIMARY KEY,
            Name VARCHAR(120)
        );`,
	"album": `
        CREATE TABLE IF NOT EXISTS Album (
            AlbumId SERIAL PRIMARY KEY,
            Title VARCHAR(160) NOT NULL,
            ArtistId INTEGER NOT NULL,
            FOREIGN KEY (ArtistId) REFERENCES Artist (ArtistId) 
        );`,

	"employee": `
        CREATE TABLE IF NOT EXISTS Employee (
            EmployeeId SERIAL PRIMARY KEY,
            LastName VARCHAR(20) NOT NULL,
            FirstName VARCHAR(20) NOT NULL,
            Title VARCHAR(30),
            ReportsTo INTEGER,
            BirthDate TIMESTAMP,
            HireDate TIMESTAMP,
            Address VARCHAR(70),
            City VARCHAR(40),
            State VARCHAR(40),
            Country VARCHAR(40),
            PostalCode VARCHAR(10),
            Phone VARCHAR(24),
            Fax VARCHAR(24),
            Email VARCHAR(60),
            FOREIGN KEY (ReportsTo) REFERENCES Employee (EmployeeId) 
        );`,
	"invoiceLine": `
        CREATE TABLE IF NOT EXISTS InvoiceLine (
            InvoiceLineId SERIAL PRIMARY KEY,
            InvoiceId INTEGER NOT NULL,
            TrackId INTEGER NOT NULL,
            UnitPrice NUMERIC(10, 2) NOT NULL,
            Quantity INTEGER NOT NULL,
            FOREIGN KEY (InvoiceId) REFERENCES Invoice (InvoiceId) ,
            FOREIGN KEY (TrackId) REFERENCES "Track" (TrackId)
        );`,
	"customer": `
        CREATE TABLE IF NOT EXISTS Customer (
            CustomerId SERIAL PRIMARY KEY,
            FirstName VARCHAR(40) NOT NULL,
            LastName VARCHAR(20) NOT NULL,
            Company VARCHAR(80),
            Address VARCHAR(70),
            City VARCHAR(40),
            State VARCHAR(40),
            Country VARCHAR(40),
            PostalCode VARCHAR(10),
            Phone VARCHAR(24),
            Fax VARCHAR(24),
            Email VARCHAR(60) NOT NULL,
            SupportRepId INTEGER,
            FOREIGN KEY (SupportRepId) REFERENCES Employee (EmployeeId) 
        );`,

	"invoice": `
        CREATE TABLE IF NOT EXISTS Invoice (
            InvoiceId SERIAL PRIMARY KEY,
            CustomerId INTEGER NOT NULL,
            InvoiceDate TIMESTAMP NOT NULL,
            BillingAddress VARCHAR(70),
            BillingCity VARCHAR(40),
            BillingState VARCHAR(40),
            BillingCountry VARCHAR(40),
            BillingPostalCode VARCHAR(10),
            Total NUMERIC(10, 2) NOT NULL,
            FOREIGN KEY (CustomerId) REFERENCES Customer (CustomerId) 
        );`,

	"track": `
        CREATE TABLE IF NOT EXISTS Track (
            TrackId SERIAL PRIMARY KEY,
            Name VARCHAR(200) NOT NULL,
            AlbumId INTEGER,
            MediaTypeId INTEGER NOT NULL,
            GenreId INTEGER,
            Composer VARCHAR(220),
            Milliseconds INTEGER NOT NULL,
            Bytes INTEGER,
            UnitPrice NUMERIC(10, 2) NOT NULL,
            FOREIGN KEY (MediaTypeId) REFERENCES MediaType (MediaTypeId) ,
            FOREIGN KEY (GenreId) REFERENCES Genre (GenreId) ,
            FOREIGN KEY (AlbumId) REFERENCES Album (AlbumId) 
        );`,

	"playlistTrack": `
        CREATE TABLE IF NOT EXISTS PlaylistTrack (
            PlaylistId INTEGER NOT NULL,
            TrackId INTEGER NOT NULL,
            PRIMARY KEY (PlaylistId, TrackId),
            FOREIGN KEY (PlaylistId) REFERENCES Playlist (PlaylistId) ,
            FOREIGN KEY (TrackId) REFERENCES Track (TrackId) 
        );`,
}

var DropTableStatements = []string{
	"DROP TABLE IF EXISTS InvoiceLine CASCADE;",
	"DROP TABLE IF EXISTS PlaylistTrack CASCADE;",
	"DROP TABLE IF EXISTS Track CASCADE;",
	"DROP TABLE IF EXISTS Invoice CASCADE;",
	"DROP TABLE IF EXISTS Employee CASCADE;",
	"DROP TABLE IF EXISTS Customer CASCADE;",
	"DROP TABLE IF EXISTS Album CASCADE;",
	"DROP TABLE IF EXISTS Genre CASCADE;",
	"DROP TABLE IF EXISTS Artist CASCADE;",
	"DROP TABLE IF EXISTS Playlist CASCADE;",
	"DROP TABLE IF EXISTS MediaType CASCADE;",
}
