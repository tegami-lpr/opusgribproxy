Small GRIB proxy for OPUS v5.23.

Problem:
Because NOAA change path where GRIB files lay old version of OPUS can't download it.

Exapmple:

Full GRIB data:
Where OPUS try find files:
http://www.ftp.ncep.noaa.gov/data/nccf/com/gfs/prod/gfs.2021052600/gfs.t00z.pgrb2.1p00.f009
Where NOAA it have:
https://ftp.ncep.noaa.gov/data/nccf/com/gfs/prod/gfs.20210526/00/atmos/gfs.t00z.pgrb2.0p25.f009


Partial GRIB:
Where OPUS try find files:
http://www.nomads.ncep.noaa.gov/cgi-bin/filter_gfs_1p00.pl?file=gfs.t00z.pgrb2.1p00.f006&lev_150_mb=on&lev_200_mb=on&lev_250_mb=on&lev_300_mb=on&lev_350_mb=on&lev_400_mb=on&lev_450_mb=on&lev_500_mb=on&lev_550_mb=on&lev_600_mb=on&lev_650_mb=on&lev_700_mb=on&lev_750_mb=on&lev_800_mb=on&lev_850_mb=on&lev_900_mb=on&lev_surface=on&lev_tropopause=on&lev_low_cloud_layer=on&lev_middle_cloud_layer=on&lev_high_cloud_layer=on&var_HGT=on&var_TMP=on&var_UGRD=on&var_VGRD=on&var_TCDC=on&var_APCP=on&var_CAPE=on&var_LFTX=on&leftlon=0&rightlon=360&toplat=90&bottomlat=-90&dir=%2Fgfs.2021052700
Where NOAA it have:
https://nomads.ncep.noaa.gov/cgi-bin/filter_gfs_1p00.pl?file=gfs.t00z.pgrb2.1p00.f006&lev_150_mb=on&lev_200_mb=on&lev_250_mb=on&lev_300_mb=on&lev_350_mb=on&lev_400_mb=on&lev_450_mb=on&lev_500_mb=on&lev_550_mb=on&lev_600_mb=on&lev_650_mb=on&lev_700_mb=on&lev_750_mb=on&lev_800_mb=on&lev_850_mb=on&lev_900_mb=on&lev_surface=on&lev_tropopause=on&lev_low_cloud_layer=on&lev_middle_cloud_layer=on&lev_high_cloud_layer=on&var_HGT=on&var_TMP=on&var_UGRD=on&var_VGRD=on&var_TCDC=on&var_APCP=on&var_CAPE=on&var_LFTX=on&leftlon=0&rightlon=360&toplat=90&bottomlat=-90&dir=%2Fgfs.20210527/00/atmos

Solution:
Intercept OPUS requests and translate it to new path.

Drawback:
Data for GVIS is now absent on NOAA ftp server.



How to install:
Just copy exe to any place and add next line to hosts file
127.0.0.1 www.ftp.ncep.noaa.gov www.nomads.ncep.noaa.gov
