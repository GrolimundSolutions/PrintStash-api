# PrintStash
A sophisticated 3D printing filament management system that helps you track your filament inventory, optimize print settings, and manage your printing workflow.

___
---
TBD


```sql
-- Erstelle Tabellen für die Stammdaten
CREATE TABLE manufacturers (
    id SMALLINT PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE materials (
    id SMALLINT PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE colors (
    id SMALLINT PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

-- Haupttabelle für die Filament-Rollen
CREATE TABLE filament_spools (
    id SERIAL PRIMARY KEY,
    manufacturer_id SMALLINT REFERENCES manufacturers(id),
    material_id SMALLINT REFERENCES materials(id),
    color_id SMALLINT REFERENCES colors(id),
    weight_total INTEGER NOT NULL, -- Gesamtgewicht in Gramm
    weight_remaining INTEGER, -- Verbleibendes Gewicht in Gramm
    purchase_date DATE,
    price DECIMAL(10,2),
    rating SMALLINT CHECK (rating BETWEEN 1 AND 5),
    notes TEXT,
    code VARCHAR(12) GENERATED ALWAYS AS (
        LPAD(manufacturer_id::text, 2, '0') || '-' ||
        LPAD(material_id::text, 2, '0') || '-' ||
        LPAD(color_id::text, 3, '0')
    ) STORED
);

-- Tabelle für Druckeinstellungen
CREATE TABLE print_settings (
    id SERIAL PRIMARY KEY,
    filament_spool_id INTEGER REFERENCES filament_spools(id),
    nozzle_temperature INTEGER NOT NULL,
    bed_temperature INTEGER NOT NULL,
    print_speed INTEGER NOT NULL,
    retraction_distance DECIMAL(4,1),
    retraction_speed INTEGER,
    flow_rate SMALLINT DEFAULT 100,
    fan_speed SMALLINT DEFAULT 100,
    notes TEXT
);

-- Einfügen der Stammdaten
-- Hersteller
INSERT INTO manufacturers (id, name) VALUES
(1, 'eSUN'),
(2, 'Pusa'),
(3, '3DFS'),
(4, 'Anycubic'),
(5, 'Bambu'),
(6, 'Sunlu'),
(7, '3DJake'),
(8, 'PrintWithSmile');

-- Materialien
INSERT INTO materials (id, name) VALUES
(1, 'PLA'),
(2, 'PLA+'),
(3, 'ecoPLA'),
(4, 'PETG'),
(5, 'rPETG'),
(6, 'ABS'),
(7, 'ABS+'),
(8, 'ASA'),
(9, 'TPU A95');

-- Farben
INSERT INTO colors (id, name) VALUES
(1, 'Weis'),
(2, 'Schwarz'),
(3, 'Rot'),
(4, 'Grün'),
(5, 'Blau'),
(6, 'Violet'),
(7, 'Grau'),
(8, 'Gold'),
(9, 'Cold White'),
(10, 'Orange'),
(11, 'Fire Engine Red'),
(12, 'Light Blue'),
(13, 'Hellgrün'),
(14, 'Transparent'),
(15, 'Transparent Grün'),
(16, 'Transparent Gelb'),
(17, 'Transparent Violet'),
(18, 'Cloudy Grey'),
(19, 'Silver'),
(20, 'Peak Green'),
(21, 'Gelb'),
(22, 'Olive Green'),
(23, 'Solid Black'),
(24, 'Solid White'),
(25, 'Pink');

-- Beispiel für das Einfügen einer Filament-Rolle
INSERT INTO filament_spools (
    manufacturer_id, material_id, color_id, 
    weight_total, weight_remaining, 
    purchase_date, price, rating
) VALUES (
    1, 2, 2,  -- eSUN, PLA+, Schwarz
    1000, 1000,  -- 1kg Rolle, noch voll
    CURRENT_DATE, 25.99, 5
);

-- Beispiel für Druckeinstellungen
INSERT INTO print_settings (
    filament_spool_id, nozzle_temperature, bed_temperature,
    print_speed, retraction_distance, retraction_speed,
    flow_rate, fan_speed, notes
) VALUES (
    1, 210, 60,  -- Temperaturen
    50, 5.0, 45,  -- Speed und Retraction
    100, 100,  -- Flow und Fan
    'Beste Ergebnisse mit diesen Einstellungen'
);
```
