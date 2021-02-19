# rsteamsupportstatus
Generates a static page with Open Graph tags that contains Steam Support information. Useful for Discord embeds.

## Install
```
GO111MODULE=on go get -u get.cutie.cafe/rsteamsupportstatus
$GOPATH/bin/rsteamsupportstatus #output to status.html
$GOPATH/bin/rsteamsupportstatus -output=/path/to/another/status.html #output to custom location
```

Statistics information updates daily at midnight Pacific, so you'll probably want to execute this then.

## License
```
rsteamsupportstatus
Copyright (C) 2021 Alexandra Frock

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
```
