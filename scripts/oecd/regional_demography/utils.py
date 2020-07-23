# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# Flatten the multi-index dataframe to be single-index.
def multi_index_to_single_index(df):
    columns = []
    for column in df.columns:
        column = list(column)
        column[1] = str(column[1])
        columns.append(''.join(column))

    df.columns = columns
    return df.reset_index()


# Generate the dcid info.
def generate_geo_id(row, nuts, name2dcid):
    if str(row['TL']) == '1':
        # print("TL1: " + "Country/" + row['REG_ID'])
        return "Country/" + row['REG_ID']
    elif row['REG_ID'] in nuts.keys():
        # print("NUTS: " + "dcid:nuts/" + row['REG_ID'])
        return "dcid:nuts/" + row['REG_ID']
    else:
        # print("RESOLVER: " + name2dcid[row['Region']])
        return name2dcid[row['Region']]
